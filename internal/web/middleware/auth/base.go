package auth

import (
	"bfw/internal/logger"

	"github.com/gin-gonic/gin"
)

const (
	AuthNone AuthMethod = 0
	AuthJwt  AuthMethod = 1
)

type AuthMethod int

type AuthInterface interface {
	Init(map[string]interface{}) error
	AuthMiddleware() gin.HandlerFunc
	AuthAdmin() gin.HandlerFunc
	GenerateToken(string, uint, []uint, []uint) (string, error)
}

type BaseAuth struct {
}

func (*BaseAuth) Init(map[string]interface{}) error {
	logger.Warn("base auth interface does absolutely nothing")
	return nil
}

func (*BaseAuth) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (*BaseAuth) AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (*BaseAuth) GenerateToken(string, uint, []uint, []uint) (string, error) {
	logger.Warn("base auth interface generates nil token")
	return "", nil
}
