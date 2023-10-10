package utils

import (
	"bfw/internal/json"
	"bfw/internal/logger"
	"bfw/internal/web"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type PaginateOrder struct {
	Field string `json:"field"`
	Sort  string `json:"sort"`
}
type Paginate struct {
	PageSize   int              `json:"pageSize"`
	PageNumber int              `json:"pageNumber"`
	PageOrder  []*PaginateOrder `json:"pageOrder"`
}
type PaginateAndSearchParam struct {
	PaginateParam *Paginate   `json:"paginate"`
	SearchParam   interface{} `json:"search"`
}

func UnmarshalRequestData(c *gin.Context, obj interface{}) error {
	//data, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	logger.Errorf("failed to read request data: %v", err)
	//	SendResp(c, web.ECGenCorruptBody, web.EMGenCorruptBody)
	//	return err
	//}
	//if err = json.JsonUnmarshal(data, obj); err != nil {
	//	logger.Errorf("error parsing request data: %v", err)
	//	SendResp(c, web.ECGenIncorrectBody, web.EMGenIncorrectBody)
	//	return err
	//}

	if err := c.BindJSON(obj); err != nil {
		logger.Errorf("error bind request data: %v", err)
		SendBadRequestResp(c)
		return err
	}

	return nil
}

func UnmarshalRequestDataWithParams(c *gin.Context, obj interface{}) error {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("failed to read request data: %v", err)
		return err
	}
	if err = json.Unmarshal(data, obj); err != nil {
		logger.Errorf("error parsing request data: %v", err)
		return err
	}

	//BindJSON don't allow the case that json is nil
	//if err := c.BindJSON(obj); err != nil {
	//	logger.Errorf("error bind request data: %v", err)
	//	return err
	//}
	return nil
}

// UnmarshalFormDataExcelFile
// csv and excel
func UnmarshalFormDataExcelFile(c *gin.Context) (*excelize.File, error) {
	var (
		byteData []byte
		err      error
	)

	//fileHeader, err := c.FormFile("file")
	//if err != nil {
	//	return nil, err
	//}
	////convert file format
	//if common.GetFileNameSuffix(fileHeader.Filename) == "csv" {
	//	byteData = make([]byte, fileHeader.Size)
	//	byteData, err = CsvToXlsx(byteData)
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	byteData, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("error read request body: %v", err.Error())
		return nil, err
	}

	reader := bytes.NewReader(byteData)

	file, err := excelize.OpenReader(reader)
	if err != nil {
		logger.Errorf("error in excelize open reader: %v", err.Error())
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			logger.Errorf("file close error: %v", err.Error())
		}
	}()

	return file, nil
}

func SendResp(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, &web.BaseResponse{
		Code:    code,
		Message: msg,
	})
}

func SendOKResp(c *gin.Context) {
	SendResp(c, web.ECOK, web.EMOK)
}

func SendDataResp(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, &web.BaseResponseWithData{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func SendDataOKResp(c *gin.Context, data interface{}) {
	SendDataResp(c, web.ECOK, web.EMOK, data)
}

func SendListResp(c *gin.Context, code int, msg string, data interface{}, total int64) {
	c.JSON(http.StatusOK, &web.BaseResponseWithListAndTotal{
		Code:    code,
		Message: msg,
		List:    data,
		Total:   total,
	})
}

func SendListOKResp(c *gin.Context, data interface{}, total int64) {
	SendListResp(c, web.ECOK, web.EMOK, data, total)
}

func SendFileDataOKResp(c *gin.Context, fileName string, fileObject io.Reader) {
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.DataFromReader(200, -1, "application/octet-stream", fileObject, nil)
}

func SendUnknownResp(c *gin.Context) {
	c.JSON(http.StatusOK, &web.BaseResponse{
		Code:    web.ECGenUnknown,
		Message: web.EMGenUnknown,
	})
}

func SendUnknownErrorResp(c *gin.Context, error string) {
	SendResp(c, web.ECGenUnknown, error)
}

func SendErrorRequestResp(c *gin.Context) {
	c.JSON(http.StatusOK, &web.BaseResponse{
		Code:    web.ECGenCorruptBody,
		Message: web.EMGenCorruptBody,
	})
}

func SendBadRequestResp(c *gin.Context) {
	c.JSON(http.StatusOK, &web.BaseResponse{
		Code:    web.ECGenIncorrectBody,
		Message: web.EMGenIncorrectBody,
	})
}

func SendNameDuplicatedResp(c *gin.Context) {
	c.JSON(http.StatusOK, &web.BaseResponse{
		Code:    web.ECGenNameDuplicated,
		Message: web.EMGenNameDuplicated,
	})
}

func SendNoPermissionResp(c *gin.Context) {
	c.JSON(http.StatusOK, &web.BaseResponse{
		Code:    web.ECUserInsufficientPermission,
		Message: web.EMUserInsufficientPermission,
	})
}

func SendNoResourceResp(c *gin.Context) {
	c.JSON(http.StatusOK, &web.BaseResponse{
		Code:    web.ECResourceNotFound,
		Message: web.EMResourceNotFound,
	})
}
