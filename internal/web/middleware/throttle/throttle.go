package throttle

import (
	"bfw/internal/logger"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	_urls    map[string]bool = make(map[string]bool)
	_limiter *rate.Limiter
)

func InitThrottle(urls []string, maxPerSec, maxBurst int) error {
	if urls != nil {
		for _, url := range urls {
			_urls[url] = true
		}
	} else {
		logger.Warn("nil throttle url list")
	}

	if maxPerSec <= 0 || maxBurst <= 0 {
		logger.Errorf("invalid throttle settings: %d,%d", maxPerSec, maxBurst)
		return errors.New("invalid throttle settings")
	}

	_limiter = rate.NewLimiter(rate.Limit(maxPerSec), maxBurst)
	return nil
}

func Throttle() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.RequestURI()
		if _, ok := _urls[url]; ok {
			if _limiter.Allow() {
				c.Next()
				return
			}

			logger.Errorf("request flooding @ url: %s", url)
			c.Error(errors.New("limit exceeded"))
			c.AbortWithStatus(http.StatusTooManyRequests)
		}
	}
}
