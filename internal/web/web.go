package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ECOK int    = 0
	EMOK string = "OK"

	ECTokenNoToken   int    = 10001
	EMTokenNoToken   string = "token not found"
	ECTokenExpire    int    = 10002
	EMTokenExpire    string = "token expired"
	ECTokenInactive  int    = 10003
	EMTokenInactive  string = "token is inactive"
	ECTokenInvalid   int    = 10004
	EMTokenInvalid   string = "invalid token"
	ECTokenMalformed int    = 10005
	EMTokenMalformed string = "malformed token"

	ECGenUnknown        int    = 20001
	EMGenUnknown        string = "unknown error"
	ECGenCorruptBody    int    = 20002
	EMGenCorruptBody    string = "corrupted request data"
	ECGenIncorrectBody  int    = 20003
	EMGenIncorrectBody  string = "incorrect request data"
	ECGenNameDuplicated int    = 20004
	EMGenNameDuplicated string = "name duplicated"

	ECUserWrongPassword          int    = 30001
	EMUserWrongPassword          string = "wrong password"
	ECUserNotExist               int    = 30002
	EMUserNotExist               string = "user not exist"
	ECUserInsufficientPermission int    = 30003
	EMUserInsufficientPermission string = "user has no permission"

	ECResourceNotFound int    = 40001
	EMResourceNotFound string = "resource not found"
)

type BaseResponse struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

type BaseResponseWithId struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
	ID      int    `json:"id"`
}

type BaseResponseWithObj struct {
	Code      int         `json:"errcode"`
	Message   string      `json:"errmsg"`
	NewObject interface{} `json:"newObject"`
}

type BaseResponseWithIdList struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
	IDList  []uint `json:"idList"`
}

type BaseResponseWithList struct {
	Code    int         `json:"errcode"`
	Message string      `json:"errmsg"`
	List    interface{} `json:"list"`
}

type BaseResponseWithListAndTotal struct {
	Code    int         `json:"errcode"`
	Message string      `json:"errmsg"`
	List    interface{} `json:"list"`
	Total   int64       `json:"total"`
}

type BaseResponseWithData struct {
	Code    int         `json:"errcode"`
	Message string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

type BatchBaseResponse struct {
	Codes    []int    `json:"errcodes"`
	Messages []string `json:"errmsgs"`
	Total    int64    `json:"total"`
	Success  int64    `json:"success"`
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers",
			"Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, Content-Disposition")
		c.Header("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, "+
				"Access-Control-Allow-Headers, Content-Type, Content-Disposition")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
