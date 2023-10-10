package fs

import (
	"bfw/internal/logger"
	"errors"
	"sync"
)

var (
	_backend BackendType = BackendUnknown
	_inst    FSInterface = nil
	_baseDir string      = ""
)

func Get() FSInterface {
	return _inst
}

func GetBaseDir() string {
	return _baseDir
}

func SetFsBackend(backend BackendType, baseDir string) error {
	_backend = backend
	_baseDir = baseDir
	var err error
	if _inst, err = ProduceFileSystem(); err != nil {
		logger.Errorf("error inner producing file system: %v", err)
		return err
	}

	return nil
}

func ProduceFileSystem() (FSInterface, error) {
	return InstantiateFileSystem(_backend)
}

func InstantiateFileSystem(backend BackendType) (FSInterface, error) {
	switch backend {
	case BackendLocal:
		return &LocalFS{
			Cache: make(map[string]*CacheFile),
			Mux:   sync.Mutex{},
		}, nil
	default:
		logger.Errorf("unimplemented file system: %d", backend)
		break
	}
	return nil, errors.New("Not implemented")
}
