package middlewares

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "http://localhost:5173") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

var ip2limiter map[string]*Limiter

type Limiter struct {
	l       *rate.Limiter
	timeout time.Time
}

func RateLimit(d time.Duration, times int) gin.HandlerFunc {
	sync.OnceFunc(func() {
		ip2limiter = make(map[string]*Limiter)
		go func() {
			t := time.NewTicker(5 * time.Second)
			for range t.C {
				now := time.Now()
				for k, l := range ip2limiter {
					if l == nil || l.timeout.After(now) {
						delete(ip2limiter, k)
					}
				}
			}
		}()
	})
	return func(c *gin.Context) {
		ip := c.ClientIP()
		// get limiter
		l, ok := ip2limiter[ip]
		if !ok {
			l = &Limiter{
				l: rate.NewLimiter(rate.Every(d), times),
			}
		}
		l.timeout = time.Now().Add(30 * time.Minute)
		if l.l.Allow() {
			c.Next()
		} else {
			// TODO log
			c.Abort()
			log.Info()
		}
	}
}
