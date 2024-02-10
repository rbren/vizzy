package llm

import (
	"github.com/rbren/vizzy/pkg/files"
)

type Client interface {
	Query(string, string) (string, error)
	Copy() Client
	SetDebugFileManager(files.FileManager)
}
