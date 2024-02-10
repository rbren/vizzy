package files

import (
	"os"
	"strings"
	"sync"
)

type FileManager interface {
	ListFilesRecursive(prefix string) ([]string, error)
	ListDirectories(prefix string) ([]string, error)
	ReadFile(key string) ([]byte, error)
	ReadJSON(key string, v interface{}) error
	WriteFile(key string, content []byte) error
	WriteJSON(key string, v interface{}) error
	CheckFileExists(key string) (bool, error)
	DeleteFile(key string) error
	DeleteRecursive(key string) error
	DeleteFileIfExists(key string) error
	CopyDirectory(sourcePrefix, destinationPrefix string) error
}

var singleton FileManager
var once sync.Once

func GetFileManager() FileManager {
	once.Do(func() {
		strategy := os.Getenv("STORAGE_STRATEGY")
		if strategy == "" {
			strategy = "local"
		}
		if strategy == "local" {
			basePath := os.Getenv("LOCAL_STORAGE_PATH")
			if basePath == "" {
				cwd, err := os.Getwd()
				if err != nil {
					panic("Could not get current working directory")
				}
				basePath = cwd + "/storage/"
			}
			if !strings.HasSuffix(basePath, "/") {
				basePath += "/"
			}
			singleton = LocalFileManager{
				BasePath: basePath,
			}
		} else if strategy == "s3" {
			inst, err := newS3Manager()
			if err != nil {
				panic("Could not connect to aws S3")
			}
			singleton = inst
		}
	})
	return singleton
}
