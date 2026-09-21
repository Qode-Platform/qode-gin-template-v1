// A minimal Gin service shaped for the fleet: every route hangs off one
// router group so a single value moves the whole app under the ingress prefix.
package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// basePath returns the fleet's ingress prefix, normalised to "" or
// "/leading/no-trailing-slash".
//
// The fleet injects BASE_PATH as /direct/<agent>:<port> and nginx forwards that
// prefix UNCHANGED, so every route must live under it. Empty means standalone:
// serve at the host root.
func basePath() string {
	raw := strings.Trim(strings.TrimSpace(os.Getenv("BASE_PATH")), "/")
	if raw == "" {
		return ""
	}
	return "/" + raw
}

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

	// One group for the whole app. With BASE_PATH empty this is just "/".
	g := r.Group(basePath())
	g.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "gin-template", "base_path": basePath()})
	})
	return r
}

func main() {
	addr := ":" + port()
	log.Printf("gin-template listening on %s (base_path=%q)", addr, basePath())
	if err := newRouter().Run(addr); err != nil {
		log.Fatal(err)
	}
}
