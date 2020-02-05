package controller

import (
	"github.com/labstack/echo"
	"shumin-project/admin-blog-web/model"
	"shumin-project/admin-blog-web/utils"
	"strconv"
	"time"
)

func ArticlePage(ctx echo.Context) error {
	var ipt struct {
		Pi int `json:"pi"`
		Ps int `json:"ps"`
	}
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if ipt.Pi < 1 {
		ipt.Pi = 1
	}
	if ipt.Ps < 1 || ipt.Ps > 50 {
		ipt.Ps = 6
	}
	count := model.ArticleCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	mods, err := model.ArticlePage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Page("新闻数据", mods, count))
}

func ArticleDrop(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	err = model.ArticleDelete(id)
	if err != nil {
		return ctx.JSON(utils.Fail("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("删除成功"))
}

func ArticleAdd(ctx echo.Context) error {
	ipt := model.Article{}
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if ipt.Title == "" {
		return ctx.JSON(utils.ErrIpt("标题不能为空"))
	}
	ipt.Ctime = time.Now().Format("2006-01-02 15:04:05")
	ipt.Utime = ipt.Ctime
	ipt.Uid, _ = ctx.Get("uid").(int64)
	err = model.ArticleAdd(&ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("添加成功"))
}
