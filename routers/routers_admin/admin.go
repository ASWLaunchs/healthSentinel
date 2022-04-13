package routers_admin

import (
	"healthSentinel/controllers/controllers_admin"
	"healthSentinel/middlewares"

	"github.com/gin-gonic/gin"
)

func Admin(router *gin.Engine) {
	apiRouters := router.Group("/admin", middlewares.MiddlewareAuth)
	{
		//single pag.
		apiRouters.GET("/", controllers_admin.ControllersIndex{}.Index)

		//api
		apiRouters.GET("/login", controllers_admin.ControllersLogin{}.Login)
		apiRouters.GET("/loginFail", controllers_admin.ControllersLogin{}.LoginFail)
		apiRouters.GET("/userList", controllers_admin.ControllersUserList{}.UserList)
		apiRouters.GET("/export", controllers_admin.ControllersExport{}.Export)
	}
}
