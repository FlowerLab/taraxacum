package taraxacum

import (
	"github.com/FlowerLab/blackdatura"
	"github.com/FlowerLab/taraxacum/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Router() *gin.Engine {
	ret := gin.New()

	log := blackdatura.With("gin router")
	ret.Use(
		blackdatura.Ginzap(log.With(zap.String("router", "gin zap"))),
		blackdatura.RecoveryWithZap(log.With(zap.String("router", "recovery with zap"))),
	)

	ret.MaxMultipartMemory = Conf.MultipartMemory

	api := ret.Group("/api/v1")
	// todo add auth
	api.POST("/upload", uploadAction)

	static := ret.Group("/static")
	static.Use(middleware.Cache())
	static.Static("/", Conf.FilePath)

	return ret
}
