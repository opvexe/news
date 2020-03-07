package router

import (
	"github.com/labstack/echo"
	"shumin-project/admin-blog-web/controller"
)

func initAdminRouter(app *echo.Echo) {
	admin := app.Group("/admin")
	adminRouter(admin)
}

func adminRouter(admin *echo.Group) {
	admin.POST("/class/add", controller.ClassAdd)
	admin.POST("/class/edit", controller.ClassEdit)
	admin.GET("/class/drop/:id", controller.ClassDrop)

	admin.GET("/user/page", controller.UserPage)
	admin.GET("/user/drop/:id", controller.UserDrop)
	admin.POST("/user/add", controller.UserAdd)
	admin.GET("/user/get/:id", controller.UserGet)
	admin.POST("/user/edit", controller.UserEdit)
	admin.GET("/article/drop/:id", controller.ArticleDrop)

	admin.POST("/article/add", controller.ArticleAdd)
}
