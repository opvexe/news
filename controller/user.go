package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"shumin-project/admin-blog-web/model"
	"shumin-project/admin-blog-web/utils"
	"strconv"
	"time"
)

//重定向
func Index(ctx echo.Context) error {
	return ctx.Redirect(302, "/login.html")
}

//登录页面
func LoginView(ctx echo.Context) error {
	return ctx.Render(200, "login.html", nil)
}

////admin
//func AdmIndexView(ctx echo.Context) error {
//	return ctx.Render(200, "index.html", nil)
//}

//登录
func Login(ctx echo.Context) error {
	fmt.Println("111111111")
	var login struct {
		Num  string `json:"num"`  //用户名
		Pass string `json:"pass"` //密码
	}
	if err := ctx.Bind(&login); err != nil {
		return ctx.JSON(utils.ErrIpt("用户输入有误", err.Error()))
	}

	user, err := model.Login(login.Num)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("用户名错误", err.Error()))
	}
	if user.Pass != login.Pass {
		return ctx.JSON(utils.ErrIpt("用户密码输入错误"))
	}
	//生成token
	claims := model.UserClaims{
		Id:   user.Id,
		Num:  user.Num,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			//设置过期时长
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	// token 包含 h,p,sign 三部分 ；h 包含加密方式,p 加密参数,sign 签名
	secret, err := token.SignedString([]byte("Goodsense@"))
	return ctx.JSON(utils.Succ("登录成功", secret))
}

//admin 获取用户
func UserGet(ctx echo.Context) error {
	//获取id
	uid := ctx.Param("id")
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取参数失败", err.Error()))
	}
	user, err := model.GetUser(id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到该用户", err.Error()))
	}
	return ctx.JSON(utils.Succ("成功", user))
}

//admin 分页
func UserPage(ctx echo.Context) error {
	var page struct {
		Pi int `json:"pi"`
		Ps int `json:"ps"`
	}
	err := ctx.Bind(&page)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据错误", err.Error()))
	}
	if page.Ps < 1 || page.Ps > 50 {
		page.Ps = 6
	}
	if page.Pi < 1 {
		return ctx.JSON(utils.ErrIpt("输入数据有误"))
	}
	count := model.CountUser()
	if count == 0 {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	p, err := model.Page(page.Pi, page.Ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Page("用户删数据", p, count))
}

//admin 删除用户
func UserDrop(ctx echo.Context) error {
	//获取id
	d := ctx.Param("id")
	id, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取参数失败", err.Error()))
	}
	//获取uid
	uid, _ := ctx.Get("uid").(int64)
	if uid == id {
		return ctx.JSON(utils.Fail("不能删除自己"))
	}
	err = model.DeleteUser(uid)
	if err != nil {
		return ctx.JSON(utils.Fail("删除失败"))
	}
	return ctx.JSON(utils.Succ("操作成功"))
}

//添加用户
func UserAdd(ctx echo.Context) error {
	var user model.User
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if user.Num == "" {
		return ctx.JSON(utils.ErrIpt("用户账号不能为空"))
	}
	if user.Name == "" {
		return ctx.JSON(utils.ErrIpt("用户名不能为空"))
	}
	if user.Pass == "" {
		return ctx.JSON(utils.ErrIpt("用户名不能为空"))
	}
	if model.ExistsUser(user.Num) {
		return ctx.JSON(utils.ErrIpt("当前账号已经存在,请重新输入"))
	}
	user.Ctime = time.Now().Format("2006-01-02 15:04:05")
	err = model.AddUser(&user)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败"))
	}
	return ctx.JSON(utils.Succ("操作成功"))
}

//编辑用户
func UserEdit(ctx echo.Context) error {
	var user model.User
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if user.Num == "" {
		return ctx.JSON(utils.ErrIpt("用户账号不能为空"))
	}
	if model.ExistsUser(user.Num) {
		return ctx.JSON(utils.ErrIpt("当前账号已经存在,请重新输入"))
	}
	err = model.AddUser(&user)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("操作成功"))
}
