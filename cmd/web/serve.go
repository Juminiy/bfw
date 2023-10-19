package web

import (
	"bfw/cmd/conf"
	"bfw/cmd/orm"
	"bfw/cmd/orm/models"
	"context"
	"errors"
	"flag"

	"bfw/cmd/web/handlers"
	"bfw/internal/cipher"
	"bfw/internal/db"
	"bfw/internal/fs"
	"bfw/internal/logger"
	"bfw/internal/web"
	"bfw/internal/web/middleware/auth"
	"bfw/internal/web/middleware/throttle"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gopkg.in/yaml.v3"
)

var (
	c                                 conf.Conf
	DatabaseMysqlInstBusinessDatabase int = 0
	DatabaseMysqlInstUserDataDatabase int = 1
)

func ServeRun() {
	// parse arguments
	enc := flag.String("enc", "", "string to encrypt")
	dec := flag.String("dec", "", "string to decrypt")
	key := flag.String("key", "", "encryption key")
	// config environment variables
	confFile := flag.String("config", "./conf/bfw.yaml", "config env")
	flag.Parse()

	initConfig(*confFile)
	initCipher(*enc, *dec, *key)
	initFileSystem()
	initDatabases()
	initMinio()
	initGin()
}

func setUrl(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", handlers.Pong)
		setModelUrl(api)
		setAdminUrl(api)
	}
}

func setModelUrl(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	if c.Web.Auth.Enable {
		v1.Use(auth.Get().AuthMiddleware())
	}

	minioOptGroup := v1.Group("/minio")
	{
		minioOptGroup.POST("/url/pre-signed/put", handlers.GetObjectPreSignedPutUrlByModuleNameAndFileName)
		minioOptGroup.POST("/url/pre-signed/get", handlers.GetObjectPreSignedGetUrlByModuleNameAndFileName)
		minioOptGroup.POST("/url/pre-signed/get/list", handlers.GetObjectPreSignedGetUrlListByModuleNameAndFileName)
	}

}

func setAdminUrl(api *gin.RouterGroup) {
	admin := api.Group("/admin")
	if c.Web.Auth.Enable {
		admin.Use(auth.Get().AuthMiddleware())
		admin.Use(auth.Get().AuthAdmin())
	}

}

func initConfig(confFile string) {
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		logger.Infof("error reading config file: %v ", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		logger.Infof("error parsing config file: %v", err)
		return
	}

	c.GlobalValueInit()
}

func initCipher(enc string, dec string, key string) {
	if len(enc) > 0 && len(key) > 0 {
		encrypted, err := cipher.Encrypt(enc, key)
		if err != nil {
			logger.Infof("error encrypting: %v\n", err)
		} else {
			logger.Infof("encrypted string: [%s]\n", encrypted)
		}
		return
	}

	if len(dec) > 0 && len(key) > 0 {
		decrypted, err := cipher.Decrypt(dec, key)
		if err != nil {
			logger.Infof("error decrypting: %v\n", err)
		} else {
			logger.Infof("encrypted string: [%s]\n", decrypted)
		}
		return
	}
	logger.InitLogger(c.Log.App.Path, c.Log.App.Level)
	logger.Info("logger initiated")

	cipher.SetGlobalKey(c.Encrypt.Code)
}

func initFileSystem() {
	logger.Infof("setting fs backend: %s", fs.BackendMap[fs.BackendType(c.FileSystem.Backend)])
	if err := fs.SetFsBackend(fs.BackendType(c.FileSystem.Backend), c.FileSystem.BaseDir); err != nil {
		logger.Errorf("error setting fs backend: %v", err)
		return
	}
	if err := fs.Get().Init(c.FileSystem.Extras); err != nil {
		logger.Errorf("error initializing fs backend: %v", err)
		return
	}
}

func initDatabases() {
	initDatabase()
}

func initDatabase() {
	logger.Infof("setting db backend: %s", db.BackendMap[db.BackendType(c.Database.Backend)])
	if err := db.SetDbBackend(db.BackendType(c.Database.Backend)); err != nil {
		logger.Errorf("error setting db backend: %v", err)
		return
	}

	content, err := decryptDatabaseMysqlConnectContent()
	if err != nil {
		return
	}

	c.Database.Username = content.Username
	c.Database.Password = content.Password
	c.Database.Database = content.Database
	c.Database.DatabaseUserData = content.DatabaseUserData

	// business database
	if err = connectDatabaseAndInitOrm(DatabaseMysqlInstBusinessDatabase); err != nil {
		return
	}
	// map models and tables
	if err = orm.InitAllTables(); err != nil {
		logger.Errorf("failed to init tables: %v", err)
		return
	}
	// callbacks
	if err = orm.InitAllCallback(); err != nil {
		logger.Errorf("failed to init callback plugin: %v", err)
		return
	}

	if c.Database.DatabaseUserData == "" {
		logger.Errorf("failed to read user data database")
		return
	}

	// change the connectDb content
	c.Database.Database = c.Database.DatabaseUserData

	// user data database
	if err = connectDatabaseAndInitOrm(DatabaseMysqlInstUserDataDatabase); err != nil {
		return
	}
}

type MysqlConnectContent struct {
	Username         string
	Password         string
	Database         string
	DatabaseUserData string
}

func decryptDatabaseMysqlConnectContent() (*MysqlConnectContent, error) {
	dbUsername, err := cipher.Decrypt(c.Database.Username, c.Encrypt.Code)
	if err != nil {
		logger.Errorf("error decrypting db username: %v", err)
		return nil, err
	}

	dbPassword, err := cipher.Decrypt(c.Database.Password, c.Encrypt.Code)
	if err != nil {
		logger.Errorf("error decrypting db password: %v", err)
		return nil, err
	}

	dbDatabase, err := cipher.Decrypt(c.Database.Database, c.Encrypt.Code)
	if err != nil {
		logger.Errorf("error decrypting db database: %v", err)
		return nil, err
	}

	dbDatabaseUserData, err := cipher.Decrypt(c.Database.DatabaseUserData, c.Encrypt.Code)
	if err != nil {
		logger.Errorf("error decrypting db database: %v", err)
		return nil, err
	}

	return &MysqlConnectContent{
		Username:         dbUsername,
		Password:         dbPassword,
		Database:         dbDatabase,
		DatabaseUserData: dbDatabaseUserData,
	}, nil
}

func connectDatabaseAndInitOrm(instId int) error {
	var err error = nil
	logger.Info("create `database` if not exists " + c.Database.Database)
	if err = orm.CreateDatabaseIfNotExists(c.Database.Driver, c.Database.Username, c.Database.Password, c.Database.Database, c.Database.Address, c.Database.Port); err != nil {
		logger.Errorf("error creating database: %v", err)
		return err
	}

	var dbInst db.DBInterface
	switch instId {
	case DatabaseMysqlInstBusinessDatabase:
		{
			dbInst = db.Get()
		}
	case DatabaseMysqlInstUserDataDatabase:
		{
			dbInst = db.GetUserData()
		}
	default:
		{
			return errors.New("unsupported database inst wheel")
		}
	}

	logger.Info("connecting database " + c.Database.Database)
	if err = dbInst.ConnectDb(&c.Database); err != nil {
		logger.Errorf("error while connecting database: %v", err)
		return err
	}

	logger.Info("initiating orm")
	if err = orm.InitOrmWithDbInst(dbInst, instId); err != nil {
		logger.Errorf("error initiating orm with db instance: %v", err)
		return err
	}

	return nil
}

func initMinio() {

	minioAccessKey, err := cipher.Decrypt(c.Minio.AccessKey, c.Encrypt.Code)
	if err != nil {
		logger.Errorf("error decrypting minio access_key: %v", err)
		return
	}
	minioSecretKey, err := cipher.Decrypt(c.Minio.SecretKey, c.Encrypt.Code)
	if err != nil {
		logger.Errorf("error decrypting minio secret_key: %v", err)
		return
	}

	models.MinioGlobalClient, err = minio.New(c.Minio.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		Secure: c.Minio.UseSSL,
	})
	if err != nil {
		logger.Errorf("error initiating minio global client: %v", err)
		return
	}

	models.MinioDefaultBucket, err = cipher.Decrypt(c.Minio.DefaultBucket, c.Encrypt.Code)
	if err != nil {
		logger.Errorf("error decrypting minio default_bucket: %v", err)
		return
	}
	models.MinioProxyPath = c.Minio.ProxyPath
	models.MinioDefaultPreSignedUrlTime = time.Duration(c.Minio.ObjectExpireTime) * time.Hour
}

func initGin() {

	if c.Web.Auth.Enable {
		auth.SetAuthMethod(auth.AuthMethod(c.Web.Auth.Method))
		auth.Get().Init(c.Web.Auth.Extras)
	}

	gin.DisableConsoleColor()
	accessFd, _ := os.Create(c.Log.Access.Path)
	gin.DefaultWriter = io.MultiWriter(accessFd)

	gin.SetMode(c.Mode)
	r := gin.Default()
	r.Use(gin.Recovery())
	if c.Web.Cors {
		r.Use(web.Cors())
	}

	if c.Web.Throttle.Enable {
		if err := throttle.InitThrottle(c.Web.Throttle.Urls, c.Web.Throttle.MaxPerSec, c.Web.Throttle.MaxBurst); err != nil {
			logger.Errorf("throttle initialization error: %v", err)
			return
		}
		r.Use(throttle.Throttle())
	}

	for _, v := range c.Web.Serves {
		logger.Infof("serving directories: %s > %s", v["k"], v["v"])
		//r.Static(v["k"], v["v"])
		r.Use(static.Serve(v["k"], static.LocalFile(v["v"], true)))
	}

	for _, v := range c.Web.Statics {
		logger.Infof("serving static urls: %s > %s", v["k"], v["v"])
		r.StaticFile(v["k"], v["v"])
	}

	setUrl(r)

	rand.Seed(time.Now().UnixNano())

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", c.Web.Bind, c.Web.Port),
		Handler:      r,
		WriteTimeout: time.Duration(c.Web.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(c.Web.ReadTimeout) * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		s := <-sigChan
		logger.Warn("system signal received, ready to quit: ", s)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	if c.Web.SSL.Enable {
		logger.Infof("running ssl on port %d", c.Web.Port)
		logger.Warn(server.ListenAndServeTLS(c.Web.SSL.CertFile, c.Web.SSL.KeyFile))
	} else {
		logger.Infof("running on port %d", c.Web.Port)
		logger.Warn(server.ListenAndServe())
	}
}
