package router

import (
	"github.com/labstack/echo"
	"shumin-project/admin-blog-web/controller"
)

func InitApiRouter(app *echo.Echo) {
	api := app.Group("/api")
	apiRouter(api)
}

func apiRouter(api *echo.Group) {
	api.POST("/login", controller.Login)
	api.GET("/class/all", controller.ClassAll)
	api.GET("/class/page", controller.ClassPage)
	api.GET("/class/get/:id", controller.ClassGet)
	api.GET("/article/page", controller.ArticlePage)
}
