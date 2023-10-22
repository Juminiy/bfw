package models

import (
	"net/url"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

const (
	MinioDefaultMaxObjectSize       int64  = (1 << 21) + (1 << 23)
	Minio3dModelBucket              string = "flex-bi-3d-model"
	MinioTopologyBucket             string = "flex-bi-topology"
	Minio3dModelBucketComponents    string = "components"
	Minio3dModelBucketComponentsImg string = "components-img"
	Minio3dModelBucketMeta          string = "meta"
	Minio3dModelBucketMetaImg       string = "meta-img"
	Minio3dModelBucketDevice        string = "device"
	Minio3dModelBucketDeviceImg     string = "device-img"
)

var (
	MinioGlobalClient             *minio.Client
	MinioDefaultBucket            string
	MinioProxyPath                string
	MinioDefaultPreSignedUrlTime  time.Duration
	MinioDefaultRemoveOptions     minio.RemoveObjectOptions
	MinioDefaultPutOptions        minio.PutObjectOptions
	MinioDefaultGetOptions        minio.GetObjectOptions
	MinioDefaultMakeBucketOptions minio.MakeBucketOptions
)

func init() {
	MinioDefaultRemoveOptions = minio.RemoveObjectOptions{GovernanceBypass: true}
	MinioDefaultPutOptions = minio.PutObjectOptions{ContentType: "application/octet-stream"}
	MinioDefaultGetOptions = minio.GetObjectOptions{}
	MinioDefaultMakeBucketOptions = minio.MakeBucketOptions{Region: "cn-north-1", ObjectLocking: false}
}

type MinioObjectDestination struct {
	ModuleName string `gorm:"-" json:"moduleName"`
	ModelName  string `gorm:"-" json:"modelName"`
	FileName   string `gorm:"-" json:"fileName"`
}

func (objectDest *MinioObjectDestination) GetPreSignedGetUrl(c *gin.Context) (string, error) {
	var (
		destFilePath   string
		destBucketName string
	)
	if len(objectDest.ModuleName) == 0 {
		destBucketName = MinioDefaultBucket
	} else {
		destBucketName = objectDest.ModuleName
	}
	if len(objectDest.ModelName) == 0 {
		destFilePath = objectDest.FileName
	} else {
		destFilePath = objectDest.ModelName + "/" + objectDest.FileName
	}
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+destFilePath+"\"")
	destUrl, err := MinioGlobalClient.PresignedGetObject(c, destBucketName,
		destFilePath, MinioDefaultPreSignedUrlTime, reqParams)
	if err != nil {
		return "", err
	}
	RemakeDestUrl(destUrl)
	return destUrl.String(), nil
}

func (objectDest *MinioObjectDestination) ValidateObjectAndGetObjectPath() (string, error) {
	var (
		destFilePath string
	)
	if len(objectDest.ModuleName) == 0 {
		objectDest.ModuleName = MinioDefaultBucket
	}
	if len(objectDest.ModelName) == 0 {
		destFilePath = objectDest.FileName
	} else {
		destFilePath = objectDest.ModelName + "/" + objectDest.FileName
	}
	return destFilePath, nil
}

func Validate3dModelBucketModelName(modelName string) bool {
	if modelName == Minio3dModelBucketComponents ||
		modelName == Minio3dModelBucketComponentsImg ||
		modelName == Minio3dModelBucketMeta ||
		modelName == Minio3dModelBucketMetaImg ||
		modelName == Minio3dModelBucketDevice ||
		modelName == Minio3dModelBucketDeviceImg {
		return true
	}
	return false
}

func ValidateBucketNameAliasModuleName(moduleName string) bool {
	if moduleName == MinioDefaultBucket ||
		moduleName == Minio3dModelBucket ||
		moduleName == MinioTopologyBucket {
		return true
	}
	return false
}

type MinioObjectDestinationList []MinioObjectDestination

func (objs MinioObjectDestinationList) ParseMinioObjectUrlListToTemporaryUrlList(c *gin.Context) ([]string, error) {
	var (
		wg          sync.WaitGroup
		err         error    = nil
		tempUrlList []string = make([]string, len(objs))
	)

	for i := 0; i < len(objs); i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			tempUrlList[i], err = objs[i].GetPreSignedGetUrl(c)
			if err != nil {
				return
			}
		}(&wg, i)
		if err != nil {
			return nil, err
		}
	}

	wg.Wait()
	return tempUrlList, err
}

func RemakeDestUrl(url *url.URL) {
	url.Scheme = ""
	url.Host = ""
	url.Path = MinioProxyPath + url.Path
}
