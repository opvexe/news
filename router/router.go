package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io"
	"shumin-project/admin-blog-web/controller"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t.templates = template.Must(template.ParseFiles("./views/login.html", "./views/index.html"))
	return t.templates.ExecuteTemplate(w, name, data)
}

//预编译模板[静态页面]
var renderer = &TemplateRenderer{templates: template.Must(template.ParseFiles("./views/login.html", "./views/index.html"))}

func Run() {
	app := echo.New()
	app.HideBanner = true   //隐藏Echo启动横幅
	app.Renderer = renderer //模板渲染

	// 允许跨域访问
	//app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//}))

	// 设置静态文件
	app.Static("/static", "static")
	app.Static("/views", "views")
	app.GET("/", controller.Index)
	app.GET("/login.html", controller.LoginView)

	//中间件
	app.Use(middleware.Recover())

	// 设置api
	initApiRouter(app)
	initAdminRouter(app)

	if err := app.Start(":8099"); err != nil {
		log.Error("启动服务失败:", err)
	}
}
