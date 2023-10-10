package utils

import (
	"bfw/cmd/orm"
	"bfw/internal/db"
	"bfw/internal/logger"
	"bfw/internal/web"
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	sqlPhraseDelimiter       = " "
	defaultOrderFieldAndSort = "`created_at` DESC"
	defaultPageSize          = 10
	defaultPageNumber        = 1
	defaultPageOffset        = 0
	defaultNonePageSize      = -1
	defaultNonePageNumber    = -1
	defaultNonePageOffset    = -1
	undefinedString          = ""
	UndefinedString          = ""
)

func ParseInterfaceStructQueryParam(infStruct interface{}) (string, error) {

	switch infStruct.(type) {
	// it will never be inferred to a struct{}
	case struct{}:
		{
			phrase := "(1=1"
			typeOfInfStruct := reflect.TypeOf(infStruct).Elem()
			valueOfInfStruct := reflect.ValueOf(infStruct).Elem()
			for i := 0; i < typeOfInfStruct.NumField(); i++ {
				fieldIdx := typeOfInfStruct.Field(i)
				fieldName := fieldIdx.Name
				fieldValue := valueOfInfStruct.Field(i)
				fieldValKind, fieldValStr := fieldValue.Kind(), fieldValue.String()
				fieldOpt := undefinedString
				if strings.ToLower(fieldName) == "name" {
					fieldOpt = "LIKE"
				} else {
					fieldOpt = "="
				}
				if fieldValKind == reflect.String {
					if strings.ToLower(fieldName) == "name" {
						fieldValStr = "'%" + fieldValStr + "%'"
					} else {
						fieldValStr = "'" + fieldValStr + "'"
					}
				}
				fieldName = FieldNameCamelToSnakeAndAddBackticks(fieldName)
				if fieldName != undefinedString &&
					fieldValStr != undefinedString &&
					((fieldValKind >= 2 && fieldValKind <= 11) ||
						fieldValKind == 13 || fieldValKind == 14 ||
						fieldValKind == 24) {
					phrase += " AND " + fieldName + sqlPhraseDelimiter + fieldOpt + sqlPhraseDelimiter + fieldValStr
				}
			}
			phrase += ")"
			return phrase, nil
		}
		// it will always be a map
	case map[string]interface{}:
		{
			infMap := infStruct.(map[string]interface{})
			phrase := "(1=1"
			for searchKey, searchValue := range infMap {
				sKey, sOpt, sValType, sValStr :=
					undefinedString, undefinedString,
					undefinedString, undefinedString
				sKey = FieldNameCamelToSnakeAndAddBackticks(searchKey)
				if searchKey == "name" {
					sOpt = "LIKE"
				} else {
					sOpt = "="
				}
				sValType, sValStr = InterfaceToString(searchValue)
				if sValType == "string" {
					if searchKey == "name" {
						sValStr = "'%" + sValStr + "%'"
					} else {
						sValStr = "'" + sValStr + "'"
					}
				}
				if searchKey != undefinedString &&
					sValType != undefinedString &&
					sValStr != undefinedString &&
					sValStr != "'%%'" &&
					sValStr != "''" {
					phrase += " AND " + sKey + sqlPhraseDelimiter + sOpt + sqlPhraseDelimiter + sValStr
				}
			}
			phrase += ")"
			return phrase, nil
		}
	default:
		{
			return undefinedString, errors.New("search param is invalid")
		}
	}
}

func SetNonePaginationParamInDBContext(db *gorm.DB) *gorm.DB {
	return db.Order(undefinedString).Limit(defaultNonePageSize).Offset(defaultNonePageOffset)
}

func SetDefaultPaginationParamInDBContext(c *gin.Context) *gorm.DB {
	return orm.DB(c).Order(defaultOrderFieldAndSort).Limit(defaultPageSize).Offset(defaultPageOffset)
}

func SetDefaultPaginationParamInRawSQL() string {
	return "ORDER BY" + defaultOrderFieldAndSort + sqlPhraseDelimiter +
		"LIMIT " + strconv.Itoa(defaultPageSize) + sqlPhraseDelimiter +
		"OFFSET " + strconv.Itoa(defaultPageOffset)
}

func SetPaginationParamInDBContextOrderByFieldsIteration(db *gorm.DB, orders []*PaginateOrder) *gorm.DB {
	// case n
	if orders != nil && len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(FieldNameCamelToSnakeAndAddBackticks(order.Field) + sqlPhraseDelimiter + strings.ToUpper(order.Sort))
		}
	} else {
		db = db.Order(defaultOrderFieldAndSort)
	}

	return db
}

// SetParamsInDBContext
// paramsOpen[0] <---> paginate param
// paramsOpen[1] <---> search param
func SetParamsInDBContext(c *gin.Context, paramsOpen ...bool) *gorm.DB {
	var (
		params *PaginateAndSearchParam
		ormDb  *gorm.DB = orm.DB(c)
	)

	// parse paginate and search params
	if paramsOpen != nil {
		if err := UnmarshalRequestDataWithParams(c, &params); err != nil {
			logger.Errorf("unmarshal request data error: %v", err)
			//if params == nil || params.PaginateParam == nil {
			//	ormDb = SetDefaultPaginationParamInDBContext(c)
			//}
		}
	}

	// set paginate param
	if paramsOpen != nil && len(paramsOpen) >= 1 && paramsOpen[0] {
		if params != nil && params.PaginateParam != nil {
			ormDb = SetPaginationParamInDBContextOrderByFieldsIteration(ormDb, params.PaginateParam.PageOrder)
			if params.PaginateParam.PageSize != 0 {
				ormDb.Limit(params.PaginateParam.PageSize)
				if params.PaginateParam.PageNumber != 0 {
					ormDb.Offset((params.PaginateParam.PageNumber - 1) * params.PaginateParam.PageSize)
				} else {
					ormDb.Offset(defaultPageOffset)
				}
			} else {
				ormDb.Limit(defaultPageSize)
				if params.PaginateParam.PageNumber != 0 {
					ormDb.Offset((params.PaginateParam.PageNumber - 1) * defaultPageSize)
				} else {
					ormDb.Offset(defaultPageOffset)
				}
			}
		} else {
			ormDb = SetDefaultPaginationParamInDBContext(c)
		}
	}

	// set search param
	if paramsOpen != nil && len(paramsOpen) >= 2 && paramsOpen[1] {
		if params != nil && params.SearchParam != nil {
			phrase, err := ParseInterfaceStructQueryParam(params.SearchParam)
			if err != nil {
				logger.Errorf("search param is open in ormDb context, but search param occurs error: %s", err.Error())
			} else {
				ormDb.Where(phrase)
			}
		}
	}

	return ormDb
}

// SetParamsInRawSQL
// mentioned the raw sql return
func SetParamsInRawSQL(c *gin.Context) (string, string) {
	var (
		orderByClause string = "ORDER BY "
		limitClause   string = "LIMIT "
		offsetClause  string = "OFFSET "
		paginateParam string = undefinedString
		searchParam   string = undefinedString
		params        *PaginateAndSearchParam
	)

	// parse paginate and search params
	{
		if err := UnmarshalRequestDataWithParams(c, &params); err != nil {
			logger.Errorf("unmarshal request data error: %v", err)
			//if params == nil || params.PaginateParam == nil {
			//	paginateParam = SetDefaultPaginationParamInRawSQL()
			//}
		}
	}

	// set paginate param
	{
		if params != nil && params.PaginateParam != nil {
			if params.PaginateParam.PageOrder != nil && len(params.PaginateParam.PageOrder) > 0 {
				for _, order := range params.PaginateParam.PageOrder {
					orderByClause += order.Field + sqlPhraseDelimiter + order.Sort + sqlPhraseDelimiter
				}
			} else {
				orderByClause += defaultOrderFieldAndSort + sqlPhraseDelimiter
			}

			if params.PaginateParam.PageSize != 0 {
				limitClause += strconv.Itoa(params.PaginateParam.PageSize) + sqlPhraseDelimiter
				if params.PaginateParam.PageNumber != 0 {
					offsetClause += strconv.Itoa((params.PaginateParam.PageNumber - 1) * params.PaginateParam.PageSize)
				} else {
					offsetClause += strconv.Itoa(defaultPageOffset)
				}
			} else {
				limitClause += strconv.Itoa(defaultPageSize) + sqlPhraseDelimiter
				if params.PaginateParam.PageNumber != 0 {
					offsetClause += strconv.Itoa((params.PaginateParam.PageNumber - 1) * defaultPageSize)
				} else {
					offsetClause += strconv.Itoa(defaultPageOffset)
				}
			}
			paginateParam = orderByClause + limitClause + offsetClause
		} else {
			paginateParam = SetDefaultPaginationParamInRawSQL()
		}
	}

	// set search param
	{
		if params != nil && params.SearchParam != nil {
			sParam, err := ParseInterfaceStructQueryParam(params.SearchParam)
			if err != nil {
				logger.Errorf("search param is open in raw sql, but search param occurs error: %s", err.Error())
			}
			searchParam = sParam
		}
	}

	return paginateParam, searchParam
}

// Standard CRUD operations after Callback TenantId crud plugin correct

func AddObjectOfTenants(c *gin.Context, object db.ModelInterface, whereQuery string, whereArgs ...string) {
	if err := UnmarshalRequestData(c, object); err == nil {
		var count int64 = 0
		if whereQuery != undefinedString && len(whereArgs) != 0 {
			if res := orm.DB(c).Model(object).Where(whereQuery, whereArgs).Count(&count); res.Error != nil {
				logger.Errorf("select [%s] by [%s %v] error: %v", object.TableName(), whereQuery, whereArgs, res.Error)
				SendUnknownErrorResp(c, res.Error.Error())
				return
			}
		}
		if count == 0 {
			if res := orm.DB(c).Model(object).Create(object); res.Error != nil {
				logger.Errorf("create [%s] error: %v", object.TableName(), res.Error)
				SendUnknownErrorResp(c, res.Error.Error())
				return
			} else {
				SendDataOKResp(c, object.GetID())
				return
			}
		} else {
			logger.Errorf("request data duplicated: %s %v", whereQuery, whereArgs)
			SendBadRequestResp(c)
			return
		}

	} else {
		return
	}
}

func DeleteObjectByIdOfTenants(c *gin.Context, object db.ModelInterface) {
	objectId := c.Param("id")
	if len(objectId) == 0 {
		SendBadRequestResp(c)
		return
	}
	if res := orm.DB(c).Model(object).Delete(object, objectId); res.Error != nil {
		logger.Errorf("delete [%s] by id error: %v", object.TableName(), res.Error)
		SendUnknownErrorResp(c, res.Error.Error())
		return
	} else {
		SendOKResp(c)
		return
	}
}

func UpdateObjectByIdOfTenants(c *gin.Context, object db.ModelInterface) {
	if err := UnmarshalRequestData(c, object); err == nil {
		if res := orm.DB(c).Model(object).Updates(object); res.Error != nil {
			logger.Errorf("update [%s] by id error: %v", object.TableName(), res.Error)
			SendUnknownErrorResp(c, res.Error.Error())
			return
		} else {
			SendOKResp(c)
			return
		}

	} else {
		return
	}
}

// Standard CRUD interface end

func GetObjectByIdOfTenants(c *gin.Context, object db.ModelInterface) {
	objectId := c.Param("id")
	if len(objectId) == 0 {
		SendBadRequestResp(c)
		return
	}

	if res := orm.DB(c).Model(object).First(object, objectId); res.Error != nil {
		logger.Errorf("select [%s] by id error: %v", object.TableName(), res.Error)
		SendUnknownErrorResp(c, res.Error.Error())
		return
	} else {
		SendDataOKResp(c, object)
		return
	}
}

// createNewObjectList v is the pointer to struct, return the struct pointer array
func createNewObjectList(v interface{}, n int64) []interface{} {
	structType := reflect.TypeOf(v)

	objectList := make([]interface{}, n)

	var i int64 = 0
	for ; i < n; i++ {
		structValue := reflect.New(structType).Elem()
		//structValue.FieldByName("ID").SetUint(0)
		//structValue.FieldByName("TenantId").SetUint(0)
		//structValue.FieldByName("UserId").SetUint(0)
		objectList[i] = structValue.Interface()
	}

	return objectList
}

func GetObjectListOfTenants(c *gin.Context, object db.ModelInterface,
	fieldClause, whereQuery string, whereArgs ...string) {

	var total int64
	res := SetParamsInDBContext(c, true, true).Model(object)
	if fieldClause != undefinedString {
		res = res.Select(fieldClause)
	}
	if whereQuery != undefinedString && len(whereArgs) != 0 {
		res = res.Where(whereQuery, whereArgs)
	}

	objectList := createNewObjectList(object, total)

	// the object itself pointer {interface{} | *MyStruct} reflect to the original common is interface{} not the MyStruct
	// although the array and slice
	res = res.Find(&objectList)
	res = SetNonePaginationParamInDBContext(res).Count(&total)
	if res.Error != nil {
		logger.Errorf("select [%s] list error: %v", object.TableName(), res.Error)
		SendUnknownErrorResp(c, res.Error.Error())
	} else {
		SendListOKResp(c, objectList, total)
	}

}

//
//func CloseDataSourceConnection(source *models.DataSource) error {
//	errMsg := "close dataSource instance connection error "
//	if err := source.Close(); err != nil {
//		errMsg += err.Error()
//		logger.Errorf(errMsg)
//		return errors.New(errMsg)
//	}
//	return nil
//}

// TODO: original database connection(default use in dataSource by dataQuery) pool

// old version db_utils

func GetPagination(c *gin.Context) (int, int) {
	npg := 0
	nps := -1
	var err error
	if c != nil {
		spg := c.Query("page")
		sps := c.Query("pageSize")
		if npg, err = strconv.Atoi(spg); err != nil {
			npg = 0
		}
		if nps, err = strconv.Atoi(sps); err != nil {
			nps = -1
		}
	}

	return npg, nps
}

func CheckDuplicate(c *gin.Context, table, query string, args ...interface{}) bool {
	ok := true
	var count int64
	if res := db.Orm().Table(table).Where(query, args...).Count(&count); res.Error != nil || count > 0 {
		ok = false
		logger.Errorf("table %s duplicate check fail, count = %d: %#v", table, count, res.Error)
		if res.Error != nil {
			SendResp(c, web.ECGenIncorrectBody, res.Error.Error())
		} else if count > 0 {
			SendResp(c, web.ECGenIncorrectBody, "已存在同名内容")
		}
	}
	return ok
}

func CheckDuplicateWithError(c *gin.Context, table, query string, args ...interface{}) error {
	var count int64
	if res := orm.DB(c).Table(table).Where(query, args...).Count(&count); res.Error != nil || count > 0 {
		logger.Errorf("table %s duplicate check fail, count = %d: %#v", table, count, res.Error)
		if res.Error != nil {
			return res.Error
		} else {
			return errors.New("duplicated error")
		}
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Standard CRUD operations

// Crteate / Update
func CreateObject(c *gin.Context, model db.ModelInterface) bool {
	ok := false
	if err := UnmarshalRequestData(c, model); err == nil {
		if model.GetID() <= 0 || CheckDuplicate(c, model.TableName(), "id <> ? AND name = ?", model.GetID(), model.GetName()) {
			if res := db.Orm().Create(model); res.Error != nil {
				logger.Errorf("failed to create object: %#v", res.Error)
				SendResp(c, web.ECGenUnknown, res.Error.Error())
			} else {
				ok = true
			}
		}
	} else {
		SendResp(c, web.ECGenIncorrectBody, err.Error())
	}
	return ok
}

// Get By ID
func GetObjectById(c *gin.Context, id uint, obj interface{}) bool {
	ok := false
	if res := db.Orm().Where("`id` = ?", id).First(obj); res.RowsAffected == 0 {
		SendResp(c, web.ECResourceNotFound, web.EMResourceNotFound)
	} else {
		ok = true
	}
	return ok
}

func GetObjectByStr(c *gin.Context, where string, obj interface{}) bool {
	ok := false
	if res := db.Orm().Where(where).First(obj); res.RowsAffected == 0 {
		SendResp(c, web.ECResourceNotFound, web.EMResourceNotFound)
	} else {
		ok = true
	}
	return ok
}

// List
func ListObjects(c *gin.Context, fields, where string, objects interface{}) bool {
	ok := false
	page, pageSize := GetPagination(c)
	if res := db.Orm().Select(fields).Where(where).Limit(pageSize).Offset(page * pageSize).Order("update_at DESC").Find(objects); res.Error != nil {
		SendResp(c, web.ECGenUnknown, res.Error.Error())
	} else {
		ok = true
	}
	return ok
}

// Delete
func DeleteObjectById(c *gin.Context, obj interface{}) bool {
	ok := false
	id := c.Param("id")
	if len(id) > 0 {
		if res := db.Orm().Delete(obj, id); res.Error != nil {
			logger.Errorf("failed to delete object by id %s: %#v", id, res.Error)
			SendResp(c, web.ECGenUnknown, res.Error.Error())
		} else {
			ok = true
		}
	} else {
		SendResp(c, web.ECGenIncorrectBody, web.EMGenIncorrectBody)
	}
	return ok
}
