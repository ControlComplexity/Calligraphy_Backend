package api

import (
	services "github.com/ControlComplexity/Calligraphy_Backend/service"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"github.com/mlogclub/simple/web/params"
)

/*
ChileduController 少儿书法教育
*/
type ChileduController struct {
	Ctx iris.Context
}

// GetChiledu 文章列表
func (c *ChileduController) GetChiledu() *web.JsonResult {
	essays, paging := services.ChileduService.Find(params.NewQueryParams(c.Ctx).
		EqByReq("contentdynasty").EqByReq("writerdynasty").EqByReq("calligrapher").
		EqByReq("type").LikeByReq("title").PageByReq().Desc("id"))
	return web.JsonPageData(essays, paging)
}

// GetBy 根据ID获取文章
func (c *ChileduController) GetBy(id int64) *web.JsonResult {
	essays := services.ChileduService.Get(id)
	return web.JsonPageData(essays, nil)
}

// GetRelated 相关文章列表
func (c *ChileduController) GetRelated() *web.JsonResult {
	essays, paging := services.ChileduService.FindRelated(params.NewQueryParams(c.Ctx).
		EqByReq("type").PageByReq().Desc("id"))
	return web.JsonPageData(essays, paging)
}

// GetRanked 文章排行榜
func (c *ChileduController) GetRanked() *web.JsonResult {
	essays, paging := services.ChileduService.FindRanked(params.NewQueryParams(c.Ctx).PageByReq().Desc("hits"))
	return web.JsonPageData(essays, paging)
}

// GetHit 点击量加1
func (c *ChileduController) GetHit() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonError(err)
	}
	e := services.ChileduService.AddHit(id)
	if e != nil {
		return web.JsonError(e)
	}
	return web.JsonSuccess()
}

// GetLike 喜爱次数加1
func (c *ChileduController) GetLike() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonError(err)
	}
	e := services.ChileduService.AddLike(id)
	if e != nil {
		return web.JsonError(e)
	}
	return web.JsonSuccess()
}
