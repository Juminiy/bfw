package auth

import (
	"bfw/internal/cipher"
	"bfw/internal/logger"
	"bfw/internal/web"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

var (
	ErrTokenExpired      error = errors.New("TOKEN EXPIRED")
	ErrTokenNotValidYet  error = errors.New("TOKEN INACTIVE")
	ErrTokenMalformed    error = errors.New("MALFORMED TOKEN")
	ErrTokenInvalid      error = errors.New("INVALID TOKEN")
	ErrTokenInsufficient error = errors.New("INSUFFICIENT TOKEN")

	_signKey         []byte        = nil
	_issuer          string        = ""
	_expireInMinutes time.Duration = 0
)

type JwtAuthClaims struct {
	jwt.StandardClaims
	Username  string `json:"username"`
	UserId    uint   `json:"userId"`
	RoleIds   []uint `json:"roleIds"`
	TenantIds []uint `json:"tenantIds"`
}

type JwtAuth struct {
}

func (*JwtAuth) Init(args map[string]interface{}) error {
	if val, ok := args["sign_key"]; ok {
		if decrypted, err := cipher.DefaultDecrypt(val.(string)); err != nil {
			logger.Error("failed to decrypt jwt sign key: %v", err)
			return errors.New("sign key decryption failed")
		} else {
			_signKey = []byte(decrypted)
		}
	} else {
		logger.Errorf("no jwt sign key config")
		return errors.New("no jwt sign key")
	}

	if val, ok := args["issuer"]; ok {
		var err error
		if _issuer, err = cipher.DefaultDecrypt(val.(string)); err != nil {
			logger.Error("failed to decrypt jwt issuer: %v", err)
			return errors.New("issuer decryption failed")
		}
	} else {
		logger.Errorf("no jwt issuer config")
		return errors.New("no jwt issuer")
	}

	if val, ok := args["expire"]; ok {
		_expireInMinutes = time.Duration(val.(int))
	} else {
		logger.Errorf("no jwt expire config")
		return errors.New("no jwt expire")
	}

	return nil
}

func (*JwtAuth) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusForbidden, &web.BaseResponse{
				Code:    web.ECTokenNoToken,
				Message: web.EMTokenNoToken,
			})
			c.Abort()
			logger.Error("token has not found")
			return
		}

		claims, err := VerifyToken(token, c.Request.URL)
		if err != nil {
			switch err {
			case ErrTokenExpired:
				logger.Error("token has expired")
				c.JSON(http.StatusForbidden, &web.BaseResponse{
					Code:    web.ECTokenExpire,
					Message: web.EMTokenExpire,
				})
				c.Abort()
				return

			case ErrTokenInvalid:
				logger.Error("token is invalid")
				c.JSON(http.StatusForbidden, &web.BaseResponse{
					Code:    web.ECTokenInvalid,
					Message: web.EMTokenInvalid,
				})
				c.Abort()
				return

			case ErrTokenMalformed:
				logger.Error("token is malformed")
				c.JSON(http.StatusForbidden, &web.BaseResponse{
					Code:    web.ECTokenMalformed,
					Message: web.EMTokenMalformed,
				})
				c.Abort()
				return

			case ErrTokenNotValidYet:
				logger.Error("token is inactive")
				c.JSON(http.StatusForbidden, &web.BaseResponse{
					Code:    web.ECTokenInactive,
					Message: web.EMTokenInactive,
				})
				c.Abort()
				return

			default:
				break
			}
		}

		c.Set("tenants", claims.(*JwtAuthClaims).TenantIds)
		c.Set("user", claims.(*JwtAuthClaims).UserId)
		c.Set("claims", claims)

		logger.Debug("token verify success")
	}
}

func (*JwtAuth) AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		admin := c.Request.Header.Get("admin")
		claims, _ := VerifyToken(token, c.Request.URL)
		userId := claims.(*JwtAuthClaims).UserId
		if userId == 1 && admin == "chisato" {
			return
		} else {
			c.JSON(http.StatusForbidden, &web.BaseResponse{
				Code:    web.ECTokenInvalid,
				Message: web.EMTokenInvalid,
			})
			c.Abort()
			logger.Error("token is invalid to operate admin")
			return
		}
	}
}

func (*JwtAuth) GenerateToken(username string, userId uint, roleIds []uint, tenantId []uint) (string, error) {
	claims := JwtAuthClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Minute * _expireInMinutes).Unix()),
			Issuer:    _issuer,
		},
		username,
		userId,
		roleIds,
		tenantId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(_signKey)

	if err != nil {
		logger.Errorf("failed to generate token for user %s: %v", username, err)
		return "", err
	}

	logger.Infof("generated token for user %s: %s", username, token)

	return signed, nil
}

func VerifyToken(tokenString string, url *url.URL) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtAuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return _signKey, nil
		},
	)

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*JwtAuthClaims); ok && token.Valid {
		logger.Infof("username: %s, expire: %d", claims.Username, claims.StandardClaims.ExpiresAt)
		return claims, nil
	}

	return nil, ErrTokenInvalid
}
