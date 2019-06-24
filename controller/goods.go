package controller

import (
	"comm-filter/service"
	"comm-filter/utils"
	"errors"

	"gopkg.in/go-playground/validator.v9"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

func goodsRoutes(p router.Party) {
	// 添加列表
	p.Post("/list", newList)
	// 获取列表
	p.Get("/list", validateGetListParams, getList)
	// 查询建议
	p.Get("/suggest", getSuggest)
	// 修改单条数据
	p.Put("/{goods_id:string}", updateValidation, updateItem)
	// 删除单条数据
	p.Delete("/{goods_id:string}", delItem)
}

// 添加列表
func newList(ctx iris.Context) {
	utils := ctx.Values().Get("utils").(userutils.Utils)
	serv := ctx.Values().Get("service").(service.Service)
	data := &struct {
		List []service.Item `json:"list"`
	}{}
	ctx.ReadJSON(data)
	// 逐个检查错误
	validate := validator.New()
	for _, item := range data.List {
		err := validate.Struct(item)
		if err != nil {
			utils.HTTPError(userutils.HTTPErr{
				Status: iris.StatusNotAcceptable,
				Error:  err,
				Code:   40006,
			})
			return
		}
	}

	newList := serv.Goods.AddGoods(data.List)

	utils.HTTPSuccess(userutils.SuccBody{
		Data:    newList,
		Restful: true,
	})
}

// 更新某项
func updateItem(ctx iris.Context) {
	utils := ctx.Values().Get("utils").(userutils.Utils)
	serv := ctx.Values().Get("service").(service.Service)

	item := ctx.Values().Get("reqBody").(*service.Item)
	err := serv.Goods.UpdateItem(item)
	if err.Error != nil {
		utils.HTTPError(err)
		return
	}

	utils.HTTPSuccess(userutils.SuccBody{
		Data:    item,
		Restful: true,
	})
}

// 删除某项
func delItem(ctx iris.Context) {
	utils := ctx.Values().Get("utils").(userutils.Utils)
	serv := ctx.Values().Get("service").(service.Service)
	goodsID := ctx.Params().Get("goods_id")

	err := serv.Goods.DelItem(goodsID)
	if err.Error != nil {
		utils.HTTPError(err)
		return
	}

	utils.HTTPSuccess(userutils.SuccBody{
		Data:    iris.Map{"success": true},
		Restful: true,
	})
}

// 获取列表
func getList(ctx iris.Context) {
	utils := ctx.Values().Get("utils").(userutils.Utils)
	serv := ctx.Values().Get("service").(service.Service)

	listOptions := ctx.Values().Get("reqBody").(*service.ListOptions)
	list := serv.Goods.List(listOptions)

	utils.HTTPSuccess(userutils.SuccBody{
		Data:    list,
		Restful: true,
	})
}

// getSuggest - 获取输入建议
func getSuggest(ctx iris.Context) {
	goodsID := ctx.URLParam("goods_id")
	serv := ctx.Values().Get("service").(service.Service)
	utils := ctx.Values().Get("utils").(userutils.Utils)

	sList := serv.Goods.Suggest(goodsID)

	utils.HTTPSuccess(userutils.SuccBody{
		Data:    sList,
		Restful: true,
	})
}

// 校验
func validateGetListParams(ctx iris.Context) {
	utils := ctx.Values().Get("utils").(userutils.Utils)
	// 参数
	page := ctx.URLParamInt64Default("page", 1)
	perPage := ctx.URLParamInt64Default("per_page", 15)
	goodsID := ctx.URLParamTrim("goods_id")

	listOptions := &service.ListOptions{
		Page:    page,
		PerPage: perPage,
		GoodsID: goodsID,
	}
	// 校验器
	validate := validator.New()
	err := validate.Struct(listOptions)
	// 校验错误
	if err != nil {
		utils.HTTPError(userutils.HTTPErr{
			Status: iris.StatusNotAcceptable,
			Error:  err,
			Code:   40006,
		})
		return
	}

	ctx.Values().Set("reqBody", listOptions)
	ctx.Next()
}

// func validateNewListParams(ctx iris.Context) {
// 	utils := ctx.Values().Get("utils").(userutils.Utils)
// }

func updateValidation(ctx iris.Context) {
	utils := ctx.Values().Get("utils").(userutils.Utils)
	goodsID := ctx.Params().Get("goods_id")
	postData := &struct {
		Item service.Item `json:"item"`
	}{}
	err := ctx.ReadJSON(postData)
	if err != nil {
		utils.HTTPError(userutils.HTTPErr{
			Status: iris.StatusNotAcceptable,
			Error:  err,
			Code:   40006,
		})
		return
	}

	item := &postData.Item

	if item.GoodsID == "" {
		utils.HTTPError(userutils.HTTPErr{
			Status: iris.StatusNotAcceptable,
			Error:  errors.New("参数错误"),
			Code:   40006,
		})
		return
	}
	item.GoodsID = goodsID
	ctx.Values().Set("reqBody", item)
	ctx.Next()
}
