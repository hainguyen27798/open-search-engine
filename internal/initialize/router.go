package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/open-search-engine/internal/router"
	"os"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if os.Getenv("MODE") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}

	mainRouter := r.Group("/v1")
	{
		router.AppRouter.Products.InitProductRouter(mainRouter)
	}

	return r
}
