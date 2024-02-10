package files

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
)

// S3Manager struct holds the S3 client and bucket name.
type S3Manager struct {
	Client     *s3.Client
	BucketName string
}

func newS3Manager() (S3Manager, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		return S3Manager{}, fmt.Errorf("unable to load SDK config, %w", err)
	}

	return S3Manager{
		Client:     s3.NewFromConfig(cfg),
		BucketName: os.Getenv("S3_BUCKET"),
	}, nil
}

func (m S3Manager) ListFilesRecursive(prefix string) ([]string, error) {
	listResult, err := m.Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &m.BucketName,
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}
	files := []string{}
	for _, item := range listResult.Contents {
		files = append(files, *item.Key)
	}
	return files, nil
}

// ListDirectories lists the files in the specified S3 bucket.
func (m S3Manager) ListDirectories(prefix string) ([]string, error) {
	resp, err := m.Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:    &m.BucketName,
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to list items in bucket %q, %v", m.BucketName, err)
	}

	var items []string
	for _, item := range resp.CommonPrefixes {
		name := strings.TrimPrefix(*item.Prefix, prefix)
		name = strings.TrimSuffix(name, "/")
		items = append(items, name)
	}

	return items, nil
}

// ReadFile reads the content of a file from the S3 bucket.
func (m S3Manager) ReadFile(key string) ([]byte, error) {
	logrus.Infof("ReadFile: %s", key)
	output, err := m.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &m.BucketName,
		Key:    &key,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read file from S3: %v", err)
	}
	defer output.Body.Close()

	return ioutil.ReadAll(output.Body)
}

func (m S3Manager) ReadJSON(key string, v interface{}) error {
	str, err := m.ReadFile(key)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(str, v); err != nil {
		return err
	}
	return nil
}

// WriteFile writes content to a file in the S3 bucket.
func (m S3Manager) WriteFile(key string, content []byte) error {
	_, err := m.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &m.BucketName,
		Key:    &key,
		Body:   bytes.NewReader(content),
	})
	if err != nil {
		return fmt.Errorf("unable to write file to S3: %v", err)
	}
	return nil
}

// CheckFileExists checks if a file exists in the S3 bucket.
func (m S3Manager) CheckFileExists(key string) (bool, error) {
	_, err := m.Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: &m.BucketName,
		Key:    &key,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// WriteJSON writes content to a file in the S3 bucket.
func (m S3Manager) WriteJSON(key string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return m.WriteFile(key, b)
}

func (m S3Manager) DeleteFile(key string) error {
	_, err := m.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &m.BucketName,
		Key:    &key,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m S3Manager) DeleteRecursive(key string) error {
	requiredPrefix := "projects/"
	if !strings.HasPrefix(key, requiredPrefix) || len(key) <= len(requiredPrefix) || !strings.HasSuffix(key, "/") {
		return fmt.Errorf("key %s must start with %s and include a project ID", key, requiredPrefix)
	}
	files, err := m.ListFilesRecursive(key)
	if err != nil {
		return err
	}
	logrus.Infof("Deleting %d files from S3", len(files))
	for _, file := range files {
		err := m.DeleteFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteFileIfExists checks if a file exists in the S3 bucket and deletes it if it does.
func (m S3Manager) DeleteFileIfExists(key string) error {
	// First, check if the file exists
	_, err := m.Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: &m.BucketName,
		Key:    &key,
	})
	if err != nil {
		// If the file does not exist, return (could handle specific 'NotFound' error)
		return nil
	}

	// If the file exists, delete it
	_, err = m.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &m.BucketName,
		Key:    &key,
	})
	if err != nil {
		return err
	}

	return nil
}

// CopyDirectory copies all files from one directory to another.
func (m S3Manager) CopyDirectory(sourcePrefix, destinationPrefix string) error {
	logrus.Infof("Copying directory %s to %s", sourcePrefix, destinationPrefix)
	paginator := s3.NewListObjectsV2Paginator(m.Client, &s3.ListObjectsV2Input{
		Bucket: &m.BucketName,
		Prefix: aws.String(sourcePrefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			logrus.Errorf("Failed to get page, %v", err)
			return err
		}

		for _, object := range page.Contents {
			sourceKey := *object.Key
			destinationKey := destinationPrefix + strings.TrimPrefix(sourceKey, sourcePrefix)

			copySource := fmt.Sprintf("%s/%s", m.BucketName, sourceKey)

			_, err := m.Client.CopyObject(context.TODO(), &s3.CopyObjectInput{
				Bucket:     &m.BucketName,
				CopySource: aws.String(copySource),
				Key:        aws.String(destinationKey),
			})
			if err != nil {
				logrus.Errorf("Failed to copy object %s to %s, %v", sourceKey, destinationKey, err)
				return err
			}

			logrus.Infof("Successfully copied %s to %s", sourceKey, destinationKey)
		}
	}
	return nil
}
