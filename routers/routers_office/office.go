package routers_office

import (
	"healthSentinel/controllers/controllers_office"

	"github.com/gin-gonic/gin"
)

func Office(router *gin.Engine) {
	apiRouters := router.Group("/")
	{
		apiRouters.GET("/", controllers_office.ControllersIndex{}.Index)
		apiRouters.POST("/upload", controllers_office.ControllersUpload{}.Upload)
		apiRouters.GET("/check", controllers_office.ControllersCheck{}.Check)
	}
}
