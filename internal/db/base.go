package db

import (
	"bfw/cmd/conf"
	"reflect"

	"gorm.io/gorm"
)

type BackendType int

const (
	BackendUnknown BackendType = 0
	BackendSqlite  BackendType = 1
	BackendMysql   BackendType = 2
	BackendMaria   BackendType = 3
)

var BackendMap = map[BackendType]string{
	BackendSqlite: "sqlite",
	BackendMysql:  "mysql",
	BackendMaria:  "maria",
}

type GeneratorFunc func() interface{}

type ModelInterface interface {
	TableName() string
	GetID() uint
	GetName() string
}

type StatementInterface interface {
	Execute(args ...interface{}) ([]map[string]interface{}, error)
	Close() error
}

type DBInterface interface {
	Connect(string, string, string, string, map[string]string, map[string]string) error
	ConnectDb(*conf.DatabaseConf) error
	Disconnect(params interface{}) error
	PrepareStatement(query string) (StatementInterface, error)
	Execute(query string, args ...interface{}) ([]map[string]interface{}, error)
	ExecuteType(resultType reflect.Type, query string, args ...interface{}) ([]interface{}, error)
	ExecuteStruct(generator GeneratorFunc, query string, args ...interface{}) ([]interface{}, error)

	Create(ModelInterface) error
	Update(ModelInterface) error
	Delete(ModelInterface) error

	Orm() *gorm.DB
}

/*
func (*DbBase) connect(params interface{}) (db *DbBase, err error) {
	log.Printf("NOT IMPLEMENTED")
	return nil, errors.New("NOT IMPLEMENTED DBBASE connect")
}

func (*DbBase) disconnect(params interface{}) (err error) {
	log.Printf("NOT IMPLEMENTED")
	return errors.New("NOT IMPLEMENTED DBBASE disconnect")
}

func (*DbBase) prepareStatement(query string) (db *DbBase, err error) {
	log.Printf("NOT IMPLEMENTED")
	return nil, errors.New("NOT IMPLEMENTED DBBASE disconnect")
}

func (*DbBase) executePreparedStatement(args ...interface{}) (results []interface{}, err error) {
	log.Printf("NOT IMPLEMENTED")
	return nil, errors.New("NOT IMPLEMENTED DBBASE disconnect")
}

func (*DbBase) execute(query string, args ...interface{}) (results []interface{}, err error) {
	log.Printf("NOT IMPLEMENTED")
	return nil, errors.New("NOT IMPLEMENTED DBBASE disconnect")
}
*/
