package fs

import (
	"time"
)

type FileType int
type BackendType int

const (
	DIR      FileType = 0
	REG_FILE FileType = 1

	BackendUnknown BackendType = 0
	BackendLocal   BackendType = 1
	BackendHDFS    BackendType = 2
	BackendCeph    BackendType = 3
)

var BackendMap = map[BackendType]string{
	BackendLocal: "local",
	BackendCeph:  "ceph",
	BackendHDFS:  "hdfs",
}

type FileStat struct {
	Path  string
	Type  FileType
	Size  int64
	MTime time.Time
}

type FS struct {
	Backend string
}

type File struct {
	Path     string
	Position int64
	Fd       interface{}
}

type FSInterface interface {
	Init(args map[string]interface{}) error
	Cleanup()

	Create(path string) (*File, error)
	Mkdir(path string) error
	Open(path string) (*File, error)
	Stat(file *File) (*FileStat, error)
	Read(file *File, start, length int64) ([]byte, error)
	ReadAll(file *File) ([]byte, error)
	Write(file *File, content []byte) error
	Flush(file *File) error
	Seek(file *File, offset int64, whence int) error
	CloseFile(file *File) error
	Delete(file *File) error
	DeleteFile(path string) error
	DeleteDir(path string) error
}

/*
func (*FS) Open(path string) (*File, error) {
	logger.Error("NOT IMPLEMENTED")
	return nil, errors.New("FS Open not implemented")
}

func (*FS) Read(file *File, start, length int64) ([]byte, error) {
	logger.Error("NOT IMPLEMENTED")
	return nil, errors.New("FS Read not implemented")
}

func (*FS) ReadAll(file *File) ([]byte, error) {
	logger.Error("NOT IMPLEMENTED")
	return nil, errors.New("FS ReadAll not implemented")
}

func (*FS) Stat(file *File) (*FileStat, error) {
	logger.Error("NOT IMPLEMENTED")
	return nil, errors.New("FS StatFile not implemented")
}

func (*FS) Write(file *File, content []byte) error {
	logger.Error("NOT IMPLEMENTED")
	return errors.New("FS Write not implemented")
}

func (*FS) Seek(file *File, offset int64, whence int) error {
	logger.Error("NOT IMPLEMENTED")
	return errors.New("FS Seek not implemented")
}

func (*FS) CloseFile(file *File) error {
	logger.Error("NOT IMPLEMENTED")
	return errors.New("FS CloseFile not implemented")
}

func (*FS) DeleteFile(path string) error {
	logger.Error("NOT IMPLEMENTED")
	return errors.New("FS DeleteFile not implemented")
}

func (*FS) DeleteDir(path string) error {
	logger.Error("NOT IMPLEMENTED")
	return errors.New("FS DeleteFile not implemented")
}
*/
