package orm

import (
	"bfw/internal/db"
	"bfw/internal/logger"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

const (
	DatabaseMysqlInstBusinessDatabase int    = 0
	DatabaseMysqlInstUserDataDatabase int    = 1
	SysInValidUserId                  uint64 = 0
	SysInValidTenantId                uint64 = 0
	SysInValidRoleId                  uint64 = 0
	SysAdminRoleId                    uint64 = 1
	SysDataAdminRoleId                uint64 = 2
)

var (
	_inst           db.DBInterface = nil
	_inst_user_data db.DBInterface = nil
)

const (
	GormMysqlTableOptionsKey   string = "gorm:table_options"
	GormMysqlTableOptionsValue string = "CHARSET=utf8mb4"
	SysInitialUserName         string = "admin"
	SysInitialUserPassword     string = "21232f297a57a5a743894a0e4a801fc3"
)

func InitAllTables() error {
	switch db.Get().(type) {
	case *db.Mysql:
		break
	default:
		return errors.New("UnImplemented database backend")
	}
	return nil
}

func InitOrmWithDbInst(inst db.DBInterface, instId int) error {
	if instId == DatabaseMysqlInstBusinessDatabase {
		if _inst != nil {
			logger.Errorf("Orm db instance can only be initiated once")
			return errors.New("orm db instance already initiated")
		}

		_inst = inst
	} else if instId == DatabaseMysqlInstUserDataDatabase {
		if _inst_user_data != nil {
			logger.Errorf("Orm db instance can only be initiated once")
			return errors.New("orm db instance already initiated")
		}

		_inst_user_data = inst
	} else {
		return errors.New("unsupported database instance id")
	}
	return nil
}

func CreateDatabaseIfNotExists(driver, username, password, database, address string, port int) error {
	db_, err := sql.Open(driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, address, port))
	if err != nil {
		logger.Errorf("error sql.Open: %#v", err)
		return err
	}
	defer func() {
		if err := db_.Close(); err != nil {
			logger.Errorf(err.Error())
		}
	}()

	_, err = db_.Exec("CREATE DATABASE IF NOT EXISTS " + "`" + database + "`;")
	if err != nil {
		logger.Errorf("error create database: %#v", err)
		return err
	}

	_, err = db_.Exec("USE " + database)
	if err != nil {
		logger.Errorf("error user database: %#v", err)
		return err
	}
	logger.Info("create database if not exists")
	return nil
}

func InitOrm(backend db.BackendType, driver string, username string, password string, database string,
	args map[string]string, pool map[string]string) error {
	if _inst != nil {
		logger.Errorf("Orm db instance can only be initiated once")
		return errors.New("orm db instance already initiated")
	}

	var err error

	_inst, err = db.InstantiateDatabase(backend)
	if err != nil {
		logger.Errorf("error instantiating db backend: %v", err)
		return err
	}

	if err = _inst.Connect(driver, username, password, database, args, pool); err != nil {
		logger.Errorf("error connecting database: %v", err)
		_inst = nil
		return err
	}

	return nil
}

func InitAllCallback() error {
	if err := db.Orm().Callback().Query().Before("gorm:query").Register("my_plugin:before_query", TenantSelectCallback); err != nil {
		logger.Errorf("failed to init tenant select callback plugin: %v", err)
		return err
	}
	if err := db.Orm().Callback().Update().Before("gorm:update").Register("my_plugin:before_update", TenantUpdateCallback); err != nil {
		logger.Errorf("failed to init tenant update callback plugin: %v", err)
		return err
	}
	if err := db.Orm().Callback().Delete().Before("gorm:delete").Register("my_plugin:before_delete", TenantDeleteCallback); err != nil {
		logger.Errorf("failed to init tenant delete callback plugin: %v", err)
		return err
	}
	if err := db.Orm().Callback().Create().Before("gorm:create").Register("my_plugin:before_create", TenantAddCallback); err != nil {
		logger.Errorf("failed to init tenant add callback plugin: %v", err)
		return err
	}
	return nil
}

// DB DbBusiness
func DB(c *gin.Context) *gorm.DB {
	if tenants, ok := c.Get("tenants"); ok {
		return db.Orm().Set("tenants", tenants)
	}
	return db.Orm()
}

// DB2 DbUserData
func DB2(c *gin.Context) *gorm.DB {
	return db.OrmUserData()
}

// TenantSelectCallback has been worked after correct.
func TenantSelectCallback(db *gorm.DB) {
	skipTenant, ok := db.Get("tenantFlag")
	flag := false
	if !ok || (ok && !skipTenant.(bool)) {
		modelType := reflect.TypeOf(db.Statement.Model).Elem()
		modelValue := reflect.ValueOf(db.Statement.Model).Elem()
		switch modelValue.Kind() {
		case reflect.Slice, reflect.Array:
			t := modelType.Elem()
			switch t.Kind() {
			case reflect.Struct:
				if _, ok := t.FieldByName("TenantId"); ok {
					flag = true
				}
			case reflect.Map:
				refMap := reflect.MakeMap(t).MapKeys()
				for _, key := range refMap {
					if key.String() == "tenant_id" {
						flag = true
						break
					}
				}
			}
		case reflect.Struct:
			if _, ok := modelType.FieldByName("TenantId"); ok {
				flag = true
			}
		case reflect.Map:
			refMap := reflect.MakeMap(modelType).MapKeys()
			for _, key := range refMap {
				if key.String() == "tenant_id" {
					flag = true
					break
				}
			}
		}

		if flag {
			tenants, _ := db.Get("tenants")
			db.Where("`tenant_id` in ?", tenants)
		}
	}
}

// TenantUpdateCallback has been worked after correct.
// maybe in user eyes, if select logic is correct, frontend shows them to user.
// user can update or delete it by id, the tenantId will never be used in delete or update method.
func TenantUpdateCallback(orm *gorm.DB) {
	var isHasTenant, tenantRight bool
	destType := reflect.TypeOf(orm.Statement.Model).Elem()
	destValue := reflect.ValueOf(orm.Statement.Model).Elem()
	switch destValue.Kind() {
	case reflect.Map:
		refMap := reflect.MakeMap(destType).MapKeys()
		for _, key := range refMap {
			if key.String() == "tenant_id" {
				isHasTenant = true
				break
			}
		}
	case reflect.Struct:
		destValue = reflect.ValueOf(orm.Statement.Dest).Elem()
		if _, ok := destType.FieldByName("TenantId"); ok {
			var id uint
			var oldTenantId int
			if _, ok := destType.FieldByName("ID"); ok {
				id = uint(destValue.FieldByName("ID").Uint())
				if id == 0 {
					orm.AddError(errors.New("the data must have id value"))
					return
				}
				where := fmt.Sprintf("id = %d", id)
				db.Orm().Set("tenantFlag", true).Select("tenant_id").Model(orm.Statement.Model).Where(where).Take(&oldTenantId)
			} else {
				orm.AddError(errors.New("the data must have id value"))
				return
			}
			tenants, _ := orm.Get("tenants")
			tenantRight, isHasTenant = judgeTenant(oldTenantId, tenants.([]uint))
			if !tenantRight {
				orm.AddError(errors.New("the user has no corresponding privileges"))
				return
			}
		}
	default:
		destType = destType.Elem()
		destValue = destValue.Elem()
		switch destValue.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < destValue.Len(); i++ {
				t := destType.Elem()
				v := destValue.Index(i)
				var oldTenantId, id int
				if _, ok := t.FieldByName("ID"); ok {
					id, _ = strconv.Atoi(v.FieldByName("ID").String())
					db.Orm().Select("tenant_id").Table(orm.Statement.Table).Where("id = ?", id).First(&oldTenantId)
				} else {
					orm.AddError(errors.New("the data must have id value"))
					return
				}
				if _, ok := t.FieldByName("TenantId"); ok {
					fieldTenantId := v.FieldByName("TenantId")
					if fieldTenantId.IsValid() {
						if fieldTenantId.IsZero() {
							orm.AddError(errors.New("the user has not belong any tenant group"))
							return
						}
						isHasTenant = true
					} else if fieldTenantId.Int() != int64(oldTenantId) {
						orm.AddError(errors.New("the user has no corresponding privileges"))
						return
					}
				}
			}
		}
	}
	if isHasTenant {
		tenants, _ := orm.Get("tenants")
		orm.Where("`tenant_id` in ?", tenants)
	}
	return
}

func TenantDeleteCallback(orm *gorm.DB) {
	var isHasTenant, tenantRight bool
	tenants, _ := orm.Get("tenants")
	destType := reflect.TypeOf(orm.Statement.Dest).Elem()
	destValue := reflect.ValueOf(orm.Statement.Dest).Elem()
	switch destValue.Kind() {
	case reflect.Map:
		destType = reflect.TypeOf(orm.Statement.Model).Elem()
		destValue = reflect.ValueOf(orm.Statement.Model).Elem()
		refMap := reflect.MakeMap(destType).MapKeys()
		for _, key := range refMap {
			if key.String() == "tenant_id" {
				isHasTenant = true
				break
			}
		}
	case reflect.Struct:
		if _, ok := destType.FieldByName("TenantId"); ok {
			tempDB := orm
			delete(tempDB.Statement.Clauses, "DELETE")
			model := reflect.New(destType).Interface()
			if res := tempDB.Debug().Find(&model); res.RowsAffected > 0 {
				modelValue := reflect.ValueOf(model)
				tenantRight, isHasTenant = judgeTenant(int(modelValue.Elem().FieldByName("TenantId").Uint()), tenants.([]uint))
				if !tenantRight {
					orm.AddError(errors.New("the user has no corresponding privileges"))
					return
				}
			} else {
				isHasTenant = true
			}
		}
	case reflect.Slice, reflect.Array:
		destType = destType.Elem()
		for i := 0; i < destValue.Len(); i++ {
			if _, ok := destType.FieldByName("TenantId"); ok {
				tempDB := orm
				delete(tempDB.Statement.Clauses, "DELETE")
				model := reflect.New(destType).Interface()
				if res := tempDB.Debug().Find(&model); res.RowsAffected == 0 {
					orm.AddError(errors.New("some data has already delete"))
					return
				}
				modelValue := reflect.ValueOf(model)
				tenantRight, isHasTenant = judgeTenant(int(modelValue.Elem().FieldByName("TenantId").Uint()), tenants.([]uint))
				if !tenantRight {
					orm.AddError(errors.New("the user has no corresponding privileges"))
					return
				}
			}
		}

	}
	if isHasTenant {
		orm.Where("`tenant_id` in ?", tenants)
	}
	return
}

// TenantAddCallback has still been worked.
// AddCallback Option should Use destination value
// AddCallback Option should omit primary key `id`
func TenantAddCallback(db *gorm.DB) {
	tenants, _ := db.Get("tenants")

	modelType := reflect.TypeOf(db.Statement.Model).Elem()
	modelValue := reflect.ValueOf(db.Statement.Model).Elem()
	switch modelValue.Kind() {
	case reflect.Slice, reflect.Array:
		t := modelType.Elem()
		switch t.Kind() {
		case reflect.Struct:
			if _, ok := t.FieldByName("TenantId"); ok {
				modelValue.Elem().FieldByName("TenantId").SetUint(uint64(tenants.([]uint)[0]))
			}
			//if _, ok := t.FieldByName("ID"); ok {
			//	modelValue.Elem().FieldByName("ID").SetUint(0)
			//	db.Statement.Omit("`id`")
			//}
		case reflect.Map:
			refMap := reflect.MakeMap(t).MapKeys()
			for _, key := range refMap {
				//if key.String() == "id" {
				//	modelValue.Elem().FieldByName("ID").SetUint(0)
				//	db.Statement.Omit("`id`")
				//}
				if key.String() == "tenant_id" {
					modelValue.Elem().FieldByName("TenantId").SetUint(uint64(tenants.([]uint)[0]))
					break
				}
			}
		}
	case reflect.Struct:
		if _, ok := modelType.FieldByName("TenantId"); ok {
			if len(tenants.([]uint)) == 0 {
				db.AddError(errors.New("the user has not belong any tenant group"))
			} else {
				modelValue.FieldByName("TenantId").SetUint(uint64(tenants.([]uint)[0]))
			}
		}
		//
		//if _, ok := modelType.FieldByName("ID"); ok {
		//	modelValue.Elem().FieldByName("ID").SetUint(0)
		//	db.Statement.Omit("`id`")
		//}
	case reflect.Map:
		refMap := reflect.MakeMap(modelType).MapKeys()
		for _, key := range refMap {
			//if key.String() == "id" {
			//	modelValue.Elem().FieldByName("ID").SetUint(0)
			//	db.Statement.Omit("`id`")
			//}
			if key.String() == "tenant_id" {
				modelValue.FieldByName("TenantId").SetUint(uint64(tenants.([]uint)[0]))
				break
			}
		}
	}
}

func judgeTenant(oldTenantId int, tenants []uint) (bool, bool) {
	var tenantRight, isHasTenant bool
	for _, tenant := range tenants {
		if int(tenant) == oldTenantId {
			tenantRight = true
			isHasTenant = true
			break
		}
	}
	return tenantRight, isHasTenant
}

//
//func TenantUpdateCallback(db *gorm.DB) {
//	flag := TenantIdReflectFlag(db)
//
//	if flag {
//		tenants, _ := db.Get("tenants")
//		db.Where("`tenant_id` in ?", tenants)
//	}
//}
//
//func TenantDeleteCallback(db *gorm.DB) {
//	flag := TenantIdReflectFlag(db)
//
//	if flag {
//		tenants, _ := db.Get("tenants")
//		db.Where("`tenant_id` in ?", tenants)
//	}
//}

//func TenantQueryCallback(db *gorm.DB) {
//	flag := TenantIdReflectFlag(db)
//
//	if flag {
//		tenants, _ := db.Get("tenants")
//		db.Where("tenant_id in ?", tenants)
//	}
//}
//
//func TenantInsertCallback(db *gorm.DB) {
//	//flag := TenantIdReflectFlag(db)
//	tenants, _ := db.Get("tenants")
//	user, _ := db.Get("user")
//	if tenantIds, ok := tenants.([]uint); ok && len(tenantIds) > 0 {
//		db.Statement.SetColumn("tenant_id", tenantIds[0])
//	}
//	if userId, ok := user.(uint); ok {
//		db.Statement.SetColumn("user_id", userId)
//	}
//}
//
//func TenantUpdateCallback(db *gorm.DB) {
//	flag := TenantIdReflectFlag(db)
//
//	if flag {
//		tenants, _ := db.Get("tenants")
//		db.Where("tenant_id in ?", tenants)
//	}
//}
//
//func TenantDeleteCallback(db *gorm.DB) {
//	flag := TenantIdReflectFlag(db)
//
//	if flag {
//		tenants, _ := db.Get("tenants")
//		db.Where("tenant_id in ?", tenants)
//	}
//}

func GetCurrentUserUserId(c *gin.Context) uint64 {
	if userId, ok := c.Get("user"); ok {
		return userId.(uint64)
	}
	return SysInValidUserId
}

func GetCurrentUserRoleId(c *gin.Context) uint64 {
	if roles, ok := c.Get("roles"); ok {
		return uint64(roles.([]uint)[0])
	} else {
		return SysInValidRoleId
	}
}

func GetCurrentUserTenantId(c *gin.Context) uint64 {
	if tenants, ok := c.Get("tenants"); ok {
		return uint64(tenants.([]uint)[0])
	} else {
		return SysInValidTenantId
	}
}

func CanManageUserData(c *gin.Context) bool {
	if roleId := GetCurrentUserRoleId(c); roleId == SysAdminRoleId ||
		roleId == SysDataAdminRoleId {
		return true
	}
	return false
}
