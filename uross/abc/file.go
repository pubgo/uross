package abc

import "io"

// FileStore abstracts the implementation file storage operations
type FileStore interface {
	CreateFile(req interface{}, file io.Reader) error
	CreateDir(req interface{}) error

	ListDir(req *interface{}) ([]interface{}, error)
	ReadFile(path string) (*interface{}, error)

	DeleteDir(path string) error
	DeleteFile(path string) error

	GetStoreType() interface{}
	Close() error
}
