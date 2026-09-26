// A minimal Gin service shaped for the fleet: it serves at the root of its own
// hostname, so routes mount directly on the engine.
package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// port is the port the fleet told us to listen on.
func port() string {
	if p := strings.TrimSpace(os.Getenv("PORT")); p != "" {
		return p
	}
	return "8080"
}

func newRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "gin-template"})
	})
	return r
}

func main() {
	addr := ":" + port()
	log.Printf("gin-template listening on %s", addr)
	if err := newRouter().Run(addr); err != nil {
		log.Fatal(err)
	}
}
