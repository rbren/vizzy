package files

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// LocalFileManager struct to handle local filesystem operations.
type LocalFileManager struct {
	BasePath string // Base directory for operations, analogous to BucketName in S3Manager.
}

// ensureDir ensures that the directory for the given file path exists.
func ensureDir(filePath string) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

func (l LocalFileManager) ListFilesRecursive(prefix string) ([]string, error) {
	var files []string
	err := filepath.Walk(l.BasePath+prefix,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				relPath, err := filepath.Rel(l.BasePath, path)
				if err != nil {
					return err
				}
				files = append(files, relPath)
			}
			return nil
		})
	return files, err
}

func (l LocalFileManager) ListDirectories(prefix string) ([]string, error) {
	var dirs []string
	fullPath := filepath.Join(l.BasePath, prefix)
	entries, err := ioutil.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}
	return dirs, nil
}

func (l LocalFileManager) ReadFile(key string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(l.BasePath, key))
}

func (l LocalFileManager) ReadJSON(key string, v interface{}) error {
	bytes, err := l.ReadFile(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, v)
}

func (l LocalFileManager) WriteFile(key string, content []byte) error {
	fullPath := filepath.Join(l.BasePath, key)
	if err := ensureDir(fullPath); err != nil {
		return err
	}
	return ioutil.WriteFile(fullPath, content, 0644)
}

func (l LocalFileManager) WriteJSON(key string, v interface{}) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return l.WriteFile(key, bytes)
}

func (l LocalFileManager) CheckFileExists(key string) (bool, error) {
	_, err := os.Stat(filepath.Join(l.BasePath, key))
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func (l LocalFileManager) DeleteFile(key string) error {
	return os.Remove(filepath.Join(l.BasePath, key))
}

func (l LocalFileManager) DeleteRecursive(key string) error {
	return os.RemoveAll(filepath.Join(l.BasePath, key))
}

func (l LocalFileManager) DeleteFileIfExists(key string) error {
	path := filepath.Join(l.BasePath, key)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil // File does not exist, no action needed
	}
	return os.Remove(path)
}

func (l LocalFileManager) CopyDirectory(sourcePrefix, destinationPrefix string) error {
	return filepath.Walk(filepath.Join(l.BasePath, sourcePrefix),
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relPath, err := filepath.Rel(filepath.Join(l.BasePath, sourcePrefix), path)
			if err != nil {
				return err
			}
			destPath := filepath.Join(l.BasePath, destinationPrefix, relPath)
			if info.IsDir() {
				return os.MkdirAll(destPath, os.ModePerm)
			}
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			return ioutil.WriteFile(destPath, content, info.Mode())
		})
}
