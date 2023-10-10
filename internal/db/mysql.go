package db

import (
	"bfw/cmd/conf"
	"errors"

	"gorm.io/driver/mysql"

	"fmt"
	"reflect"
	"strconv"
	"time"

	"bfw/internal/logger"

	"gorm.io/gorm"
)

type Mysql struct {
	db *gorm.DB
}

type MysqlStatement struct {
	db          *gorm.DB
	queryFormat string
}

func (me *Mysql) Orm() *gorm.DB {
	return me.db
}

func (me *Mysql) Connect(driver string, username string, password string, database string,
	args map[string]string, pool map[string]string) error {
	if len(driver) == 0 || len(username) == 0 || len(password) == 0 || len(database) == 0 {
		logger.Errorf("invalid connection parameters: %s, %s, %s, %s", driver, username, password, database)
		return errors.New("invalid connection parameters")
	}

	tryStr := fmt.Sprintf("%s:%s@/", username, password)
	tryDb, err := gorm.Open(mysql.Open(tryStr), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logger.Errorf("failed to connect to database: %v", err)
		return err
	}
	res := tryDb.Exec("CREATE DATABASE IF NOT EXISTS " + database + " DEFAULT CHARACTER SET utf8mb4")
	if res.Error != nil {
		logger.Errorf("failed to create database: %v", err)
		return res.Error
	}

	connStr := fmt.Sprintf("%s:%s@/%s?parseTime=True&loc=Asia%2fShanghai", username, password, database)
	if len(args) > 0 {
		for k, v := range args {
			if k == "parseTime" {
				continue
			}

			if k == "charset" && v == "utf8" {
				v = "utf8mb4"
			}

			connStr += fmt.Sprintf("&%s=%s", k, v)
		}
	}

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logger.Errorf("failed to connect to database: %v", err)
		return err
	}
	me.db = db

	sqlDb, _ := db.DB()
	if len(pool) > 0 {

		if val, ok := pool["maxOpen"]; ok {
			if n, err := strconv.Atoi(val); err == nil {
				sqlDb.SetMaxOpenConns(n)
			} else {
				logger.Errorf("invalid maxOpen value: %s", val)
			}
		}

		if val, ok := pool["maxIdle"]; ok {
			if n, err := strconv.Atoi(val); err == nil {
				sqlDb.SetMaxIdleConns(n)
			} else {
				logger.Errorf("invalid maxIdle value: %s", val)
			}
		}

		if val, ok := pool["maxLife"]; ok {
			if n, err := strconv.Atoi(val); err == nil {
				sqlDb.SetConnMaxLifetime(time.Duration(n) * time.Second)
			} else {
				logger.Errorf("invalid maxLife value: %s", val)
			}
		}
	}

	if err = sqlDb.Ping(); err != nil {
		logger.Errorf("database connection ping error: %v", err)
		defer func() {
			if err := sqlDb.Close(); err != nil {
				logger.Errorf(err.Error())
			}
		}()
		return err
	}

	logger.Info("database connected")

	return nil
}

func (me *Mysql) ConnectDb(conf *conf.DatabaseConf) error {
	if len(conf.Driver) == 0 || len(conf.Username) == 0 || len(conf.Password) == 0 || len(conf.Database) == 0 {
		logger.Errorf("invalid connection parameters: %s, %s, %s, %s", conf.Driver, conf.Username, conf.Password, conf.Database)
		return errors.New("invalid connection parameters")
	}

	tryStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/", conf.Username, conf.Password, conf.Address, conf.Port)
	tryDb, err := gorm.Open(mysql.Open(tryStr), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect to database: %v", err)
		return err
	}
	res := tryDb.Exec("CREATE DATABASE IF NOT EXISTS " + conf.Database + " DEFAULT CHARACTER SET utf8mb4")
	if res.Error != nil {
		logger.Errorf("failed to create database: %v", err)
		return res.Error
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=%s&maxAllowedPacket=%d", conf.Username, conf.Password, conf.Address, conf.Port, conf.Database, "Asia%2FShanghai", 0)
	if len(conf.Extras) > 0 {
		for k, v := range conf.Extras {
			if k == "parseTime" {
				continue
			}

			if k == "charset" && v == "utf8" {
				v = "utf8mb4"
			}

			connStr += fmt.Sprintf("&%s=%s", k, v)
		}
	}

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logger.Errorf("failed to connect to database: %v", err)
		return err
	}
	me.db = db

	sqlDb, _ := db.DB()
	if len(conf.Pool) > 0 {

		if val, ok := conf.Pool["maxOpen"]; ok {
			if n, err := strconv.Atoi(val); err == nil {
				sqlDb.SetMaxOpenConns(n)
			} else {
				logger.Errorf("invalid maxOpen value: %s", val)
			}
		}

		if val, ok := conf.Pool["maxIdle"]; ok {
			if n, err := strconv.Atoi(val); err == nil {
				sqlDb.SetMaxIdleConns(n)
			} else {
				logger.Errorf("invalid maxIdle value: %s", val)
			}
		}

		if val, ok := conf.Pool["maxLife"]; ok {
			if n, err := strconv.Atoi(val); err == nil {
				sqlDb.SetConnMaxLifetime(time.Duration(n) * time.Second)
			} else {
				logger.Errorf("invalid maxLife value: %s", val)
			}
		}
	}

	if err = sqlDb.Ping(); err != nil {
		logger.Errorf("database connection ping error: %v", err)
		defer func() {
			if err := sqlDb.Close(); err != nil {
				logger.Errorf(err.Error())
			}
		}()
		return err
	}

	logger.Info("database connected")

	return nil
}

func (me *Mysql) Disconnect(params interface{}) error {
	sqlDb, err := me.db.DB()
	if err != nil {
		logger.Errorf("error fetching db instance: %#v", err)
	} else {
		if err = sqlDb.Close(); err != nil {
			logger.Errorf("error closing database: %v", err)
		}
	}
	return err
}

func (me *Mysql) PrepareStatement(query string) (StatementInterface, error) {
	if len(query) == 0 {
		logger.Errorf("invalid prepared mysql statement: %s", query)
		return nil, errors.New("Invalid prepared statement")
	}

	return &MysqlStatement{
		queryFormat: query,
		db:          me.db,
	}, nil
}

func (me *MysqlStatement) Execute(args ...interface{}) ([]map[string]interface{}, error) {
	if me.db == nil {
		logger.Errorf("db not ready")
		return nil, errors.New("db is not ready")
	}

	if len(me.queryFormat) == 0 {
		logger.Errorf("query not ready")
		return nil, errors.New("query is not ready")
	}

	rows, err := me.db.Raw(me.queryFormat, args...).Rows()
	if err != nil {
		logger.Errorf("error executing prepared query: %v", err)
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Errorf(err.Error())
		}
	}()
	columns, err := rows.ColumnTypes()
	if err != nil {
		logger.Errorf("error getting result columns: %v", err)
		return nil, err
	}

	rowValues := make([]interface{}, len(columns))
	var results []map[string]interface{}

	for rows.Next() {
		row := map[string]interface{}{}
		for i, col := range columns {
			row[col.Name()] = reflect.New(col.ScanType()).Interface()
			rowValues[i] = row[col.Name()]
		}

		err = rows.Scan(rowValues...)
		if err != nil {
			logger.Errorf("error scanning results: %v", err)
			return nil, err
		}

		results = append(results, row)
	}

	return results, nil
}

func (*MysqlStatement) Close() error {
	logger.Infof("mysql prepared statement closed")
	return nil
}

func (me *Mysql) Execute(query string, args ...interface{}) ([]map[string]interface{}, error) {
	if me.db == nil {
		logger.Errorf("db not ready")
		return nil, errors.New("db is not ready")
	}

	rows, err := me.db.Raw(query, args...).Rows()
	if err != nil {
		logger.Errorf("error executing query: %v", err)
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Errorf(err.Error())
		}
	}()
	columns, err := rows.ColumnTypes()
	if err != nil {
		logger.Errorf("error getting result columns: %v", err)
		return nil, err
	}

	rowValues := make([]interface{}, len(columns))
	var results []map[string]interface{}

	for rows.Next() {
		row := map[string]interface{}{}
		for i, col := range columns {
			row[col.Name()] = reflect.New(col.ScanType()).Interface()
			rowValues[i] = row[col.Name()]
		}

		err = rows.Scan(rowValues...)
		if err != nil {
			logger.Errorf("error scanning results: %v", err)
			return nil, err
		}

		results = append(results, row)
	}

	return results, nil
}

func (me *Mysql) ExecuteStruct(generator GeneratorFunc, query string, args ...interface{}) ([]interface{}, error) {
	if me.db == nil {
		logger.Errorf("db is not ready")
		return nil, errors.New("db is not ready")
	}

	rows, err := me.db.Raw(query, args...).Rows()
	if err != nil {
		logger.Errorf("error executing query: %v", err)
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Errorf(err.Error())
		}
	}()
	var results []interface{}

	for rows.Next() {
		dst := generator()
		err = me.db.ScanRows(rows, dst)
		if err != nil {
			logger.Errorf("error scanning rows: %v", err)
			return nil, err
		}
		logger.Info(dst)
		results = append(results, dst)
	}

	return results, nil
}

func (me *Mysql) ExecuteType(resultType reflect.Type, query string, args ...interface{}) ([]interface{}, error) {
	if me.db == nil {
		logger.Errorf("db not ready")
		return nil, errors.New("db is not ready")
	}

	rows, err := me.db.Raw(query, args...).Rows()
	if err != nil {
		logger.Errorf("error executing query: %v", err)
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Errorf(err.Error())
		}
	}()
	var results []interface{}

	for rows.Next() {
		dst := reflect.New(resultType)
		err = me.db.ScanRows(rows, dst)
		if err != nil {
			logger.Errorf("error scanning rows: %v", err)
			return nil, err
		}
		logger.Info(dst)
		results = append(results, dst)
	}

	return results, nil
}

func (me *Mysql) Create(model ModelInterface) error {
	me.db.Set("gorm:table_options", "CHARSET=utf8mb4").Table(model.TableName()).Create(model)
	return nil
}

func (me *Mysql) Update(model ModelInterface) error {
	me.db.Set("gorm:table_options", "CHARSET=utf8mb4").Table(model.TableName()).Save(model)
	return nil
}

func (me *Mysql) Delete(model ModelInterface) error {
	me.db.Set("gorm:table_options", "CHARSET=utf8mb4").Table(model.TableName()).Delete(model)
	return nil
}
