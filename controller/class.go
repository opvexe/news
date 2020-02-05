package controller

import (
	"github.com/labstack/echo"
	"shumin-project/admin-blog-web/model"
	"shumin-project/admin-blog-web/utils"
	"strconv"
)

func ClassGet(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	mod, err := model.ClassGet(id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("分类数据", mod))
}

func ClassAll(ctx echo.Context) error {
	mods, err := model.ClassAll()
	if err != nil {
		ctx.JSON(utils.Fail("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("分类数据", mods))
}

func ClassPage(ctx echo.Context) error {
	pi, err := strconv.Atoi(ctx.FormValue("pi"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页索引错误", err.Error()))
	}
	if pi < 1 {
		return ctx.JSON(utils.ErrIpt("分页索引范围错误"))
	}
	ps, err := strconv.Atoi(ctx.FormValue("ps"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页大小错误", err.Error()))
	}
	if ps < 1 || ps > 50 {
		ps = 6
	}
	count := model.ClassCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	mods, err := model.ClassPage(pi, ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Page("分类分页数据", mods, count))
}

func ClassAdd(ctx echo.Context) error {
	ipt := &model.Class{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if ipt.Name == "" {
		return ctx.JSON(utils.ErrIpt("分类名称不能为空"))
	}
	err = model.ClassAdd(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("添加成功"))
}

func ClassEdit(ctx echo.Context) error {
	ipt := &model.Class{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if ipt.Name == "" {
		return ctx.JSON(utils.ErrIpt("分类名称不能为空"))
	}
	err = model.ClassEdit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("修改成功"))
}

func ClassDrop(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	err = model.ClassDelete(id)
	if err != nil {
		return ctx.JSON(utils.Fail("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("删除成功"))
}
