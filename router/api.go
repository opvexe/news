package router

import (
	"github.com/labstack/echo"
	"shumin-project/admin-blog-web/controller"
	middware "shumin-project/admin-blog-web/middleware"
)

func initApiRouter(app *echo.Echo) {
	api := app.Group("/api")
	apiRouter(api)
}

func apiRouter(api *echo.Group)  {
	api.POST("/login", controller.Login)
	api.Use(middware.JWT)
	api.GET("/class/all", controller.ClassAll)
	api.GET("/class/page", controller.ClassPage)
	api.GET("/class/get/:id", controller.ClassGet)
	api.GET("/article/page", controller.ArticlePage)
}
