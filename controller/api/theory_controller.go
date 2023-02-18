package api

import (
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	services "github.com/ControlComplexity/Calligraphy_Backend/service"
	"github.com/ControlComplexity/Calligraphy_Backend/utils"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"github.com/mlogclub/simple/web/params"
)

/*
TheoryController 书法教育理论，以文字为主，图片为辅
*/
type TheoryController struct {
	Ctx iris.Context
}

// GetTheory 书法教育理论列表
func (c *TheoryController) GetTheory() *web.JsonResult {
	theory, paging := services.TheoryService.Find(params.NewQueryParams(c.Ctx).
		EqByReq("type").LikeByReq("keyword").PageByReq().Desc("id"))
	res := make([]model.Theory, 0)
	for _, es := range theory{
		res = append(res, model.Theory{
			Model: model.Model{
				Id: es.Id,
			},
			Time: utils.FormatTimeToString(es.Time),
			Abstract: es.Abstract,
			Title: es.Title,
			Image: es.Image,
		})
	}
	return web.JsonPageData(res, paging)
}

// GetBy 根据ID获取文章
func (c *TheoryController) GetBy(id int64) *web.JsonResult {
	essays := services.TheoryService.Get(id)
	return web.JsonPageData(essays, nil)
}

// GetRelated 相关文章列表
func (c *TheoryController) GetRelated() *web.JsonResult {
	essays, paging := services.TheoryService.FindRelated(params.NewQueryParams(c.Ctx).
		EqByReq("type").PageByReq().Desc("id"))
	return web.JsonPageData(essays, paging)
}

// GetRanked 文章排行榜
func (c *TheoryController) GetRanked() *web.JsonResult {
	essays, paging := services.TheoryService.FindRanked(params.NewQueryParams(c.Ctx).PageByReq().Desc("hits"))
	return web.JsonPageData(essays, paging)
}

// GetHit 点击量加1
func (c *TheoryController) GetHit() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonError(err)
	}
	e := services.TheoryService.AddHit(id)
	if e != nil {
		return web.JsonError(e)
	}
	return web.JsonSuccess()
}

// GetLike 喜爱次数加1
func (c *TheoryController) GetLike() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonError(err)
	}
	e := services.TheoryService.AddLike(id)
	if e != nil {
		return web.JsonError(e)
	}
	return web.JsonSuccess()
}
