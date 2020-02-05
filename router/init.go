package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"shumin-project/admin-blog-web/controller"
	"shumin-project/admin-blog-web/model"
	"shumin-project/admin-blog-web/utils"
)

var debug = true

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if debug {
		t.templates = template.Must(template.ParseFiles("./views/login.html", "./views/index.html"))
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

//预编译模板
var renderer = &TemplateRenderer{templates: template.Must(template.ParseFiles("./views/login.html", "./views/index.html"))}

func Run() {
	app := echo.New()
	app.HideBanner = true   //隐藏启动横幅
	app.Renderer = renderer //模板渲染
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	app.Static("/static", "static")
	app.Static("/views", "views")
	app.GET("/", controller.Index)
	app.GET("/login.html", controller.LoginView)
	//api
	api := app.Group("/api")
	apiRouter(api)
	//admin
	admin := app.Group("/admin")
	adminRouter(admin)
	//启动
	if err := app.Start(":8099"); err != nil {
		fmt.Println("启动服务失败:", err)
	}
}

// ************** api路由分组 ***************** //
func apiRouter(api *echo.Group) {
	api.POST("/login", controller.Login)
	api.GET("/class/all", controller.ClassAll)
	api.GET("/class/page", controller.ClassPage)
	api.GET("/class/get/:id", controller.ClassGet)
	api.GET("/article/page", controller.ArticlePage)
}

// *************** 后台管理分组 *************** //
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

// ******************* 中间件 ******************** //
func MiddlewareJWT(ctx echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Response().Header().Set(echo.HeaderServer, "Echo/999")
		tokenString := context.FormValue("token")
		claims := model.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("Goodsense@"), nil
		})
		if err == nil && token.Valid {
			//验证通过
			context.Set("uid", claims.Id)
			return ctx(context)
		} else {
			return context.JSON(utils.ErrJwt("token验证失败"))
		}
	}
}
