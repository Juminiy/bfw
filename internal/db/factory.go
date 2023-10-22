package db

import (
	"bfw/internal/logger"
	"errors"

	"gorm.io/gorm"
)

var (
	backend      BackendType = BackendUnknown
	inst         DBInterface = nil
	instUserData DBInterface = nil
)

func Get() DBInterface {
	return inst
}

func Orm() *gorm.DB {
	if inst != nil {
		return inst.Orm()
	}
	return nil
}

func GetUserData() DBInterface {
	return instUserData
}

func OrmUserData() *gorm.DB {
	if instUserData != nil {
		return instUserData.Orm()
	}
	return nil
}

func SetDbBackend(_backend BackendType) error {
	backend = _backend
	var err error
	if inst, err = ProduceDatabase(); err != nil {
		logger.Errorf("error inner producing db backend: %v", err)
		return err
	}
	if instUserData, err = ProduceDatabase(); err != nil {
		logger.Errorf("error inner producing db backend: %v", err)
		return err
	}
	return nil
}

func ProduceDatabase() (DBInterface, error) {
	return InstantiateDatabase(backend)
}

func InstantiateDatabase(_backend BackendType) (DBInterface, error) {
	switch _backend {
	case BackendMysql:
		return &Mysql{}, nil
	default:
		logger.Errorf("unimplemented database: %d", _backend)
		break
	}
	return nil, errors.New("unimplemented database wheel ")
}
