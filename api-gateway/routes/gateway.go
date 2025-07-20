package routes

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/shehbaazsk/go-commerce-micorservices/api-gateway/config"
)

func proxy(targetHost string) gin.HandlerFunc {
	return func(c *gin.Context) {
		target, _ := url.Parse(targetHost)
		proxy := httputil.NewSingleHostReverseProxy(target)

		c.Request.URL.Scheme = target.Scheme
		c.Request.URL.Host = target.Host
		c.Request.Host = target.Host

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func SetupRoutes(r *gin.Engine, cfg config.Config) {
	v1 := r.Group("/api/v1")

	// v1.Any("/user/*path", proxy(config.Cfg.UserServiceURL))
	v1.Any("/roles/*path", proxy(cfg.UserServiceURL))
	// v1.Any("/permission/*path", proxy(config.Cfg.UserServiceURL))

}
