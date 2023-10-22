package handlers

import (
	"bfw/cmd/orm/models"
	"bfw/internal/logger"
	"bfw/internal/utils"
	"errors"
	"mime/multipart"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.String(200, "pong")
}

func GetObjectRelativeUrlByModuleNameOf3dModel(c *gin.Context) {
	moduleName := c.Param("model_name")
	if len(moduleName) == 0 || !models.Validate3dModelBucketModelName(moduleName) {
		utils.SendBadRequestResp(c)
		return
	}
	fileHeader, file := ReadFormFileToGetFileAndFileHeader(c)
	fileName := moduleName + "/" + fileHeader.Filename
	_, err := models.MinioGlobalClient.PutObject(c, models.Minio3dModelBucket,
		fileName, file, fileHeader.Size, models.MinioDefaultPutOptions)

	if err != nil {
		utils.SendUnknownErrorResp(c, err.Error())
		return
	}

	utils.SendDataOKResp(c, "./"+models.Minio3dModelBucket+"/"+fileName)
}

func ReadFormFileToGetFileAndFileHeader(c *gin.Context) (*multipart.FileHeader, multipart.File) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		utils.SendUnknownErrorResp(c, err.Error())
		return nil, nil
	}

	// file size validator
	if fileHeader.Size > models.MinioDefaultMaxObjectSize {
		utils.SendUnknownErrorResp(c, errors.New("file size is too large").Error())
		return nil, nil
	}

	file, err := fileHeader.Open()
	if err != nil {
		utils.SendUnknownErrorResp(c, err.Error())
		return nil, nil
	}
	defer func() {
		if err0 := file.Close(); err0 != nil {
			logger.Errorf(err0.Error())
		}
	}()
	return fileHeader, file
}

func GetObjectPreSignedPutUrlByModuleNameAndFileName(c *gin.Context) {
	GetPreSignedUrl(c, true)
}

func GetObjectPreSignedGetUrlByModuleNameAndFileName(c *gin.Context) {
	GetPreSignedUrl(c, false)
}

func GetPreSignedUrl(c *gin.Context, opt bool) {
	var (
		minioObjectDest models.MinioObjectDestination
		destModuleName  string = models.MinioDefaultBucket
		destModelName   string = ""
		destFileName    string = ""
		destUrl         *url.URL
		err             error = nil
	)
	// parse data
	{
		err = utils.UnmarshalRequestData(c, &minioObjectDest)
		if err != nil {
			return
		}
		if len(minioObjectDest.FileName) == 0 {
			utils.SendBadRequestResp(c)
			return
		}
		destFileName = minioObjectDest.FileName
		if len(minioObjectDest.ModuleName) != 0 {
			destModuleName = minioObjectDest.ModuleName
		}
		if len(minioObjectDest.ModelName) != 0 {
			destModelName = minioObjectDest.ModelName
			destFileName = destModelName + "/" + destFileName
		}
	}

	// create bucket if not exists
	{
		exists, err := models.MinioGlobalClient.BucketExists(c, destModuleName)
		if err != nil {
			utils.SendUnknownErrorResp(c, err.Error())
			return
		}
		if !exists {
			err = models.MinioGlobalClient.MakeBucket(c, destModuleName, models.MinioDefaultMakeBucketOptions)
			if err != nil {
				utils.SendUnknownErrorResp(c, err.Error())
				return
			}
		}
	}

	// validate bucket name
	{
		if !models.ValidateBucketNameAliasModuleName(destModuleName) {
			utils.SendBadRequestResp(c)
			return
		}
	}

	// get pre_signed url
	{
		reqParams := make(url.Values)
		reqParams.Set("response-content-disposition", "attachment; filename=\""+destFileName+"\"")
		if opt {
			destUrl, err = models.MinioGlobalClient.PresignedPutObject(c, destModuleName,
				destFileName, models.MinioDefaultPreSignedUrlTime)
		} else {
			destUrl, err = models.MinioGlobalClient.PresignedGetObject(c, destModuleName,
				destFileName, models.MinioDefaultPreSignedUrlTime, reqParams)
		}
		if err != nil {
			utils.SendUnknownErrorResp(c, err.Error())
			return
		}
	}
	models.RemakeDestUrl(destUrl)
	utils.SendDataOKResp(c, destUrl.String())
}

func GetObjectPreSignedGetUrlListByModuleNameAndFileName(c *gin.Context) {
	var (
		objs models.MinioObjectDestinationList
	)
	err := utils.UnmarshalRequestData(c, &objs)
	if err != nil {
		return
	}
	urls, err := objs.ParseMinioObjectUrlListToTemporaryUrlList(c)
	if err != nil {
		utils.SendUnknownErrorResp(c, err.Error())
		return
	}
	utils.SendDataOKResp(c, urls)
}
