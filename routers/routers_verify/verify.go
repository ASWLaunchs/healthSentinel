package routers_verify

import (
	"healthSentinel/controllers/controllers_verify"

	"github.com/gin-gonic/gin"
)

func Verify(router *gin.Engine) {
	apiRouters := router.Group("/verify")
	{
		apiRouters.POST("/admin", controllers_verify.ControllersVerify{}.Admin)
	}
}
