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
TechController 少儿书法教育
*/
type TechController struct {
	Ctx iris.Context
}

// GetTech 文章列表
func (c *TechController) GetTech() *web.JsonResult {
	essays, paging := services.TechService.Find(params.NewQueryParams(c.Ctx).
		EqByReq("contentdynasty").EqByReq("writerdynasty").EqByReq("calligrapher").
		EqByReq("type").LikeByReq("title").PageByReq().Desc("id"))
	res := make([]model.Tech, 0)
	for _, es := range essays{
		res = append(res, model.Tech{
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
func (c *TechController) GetBy(id int64) *web.JsonResult {
	essays := services.TechService.Get(id)
	return web.JsonPageData(essays, nil)
}

// GetRelated 相关文章列表
func (c *TechController) GetRelated() *web.JsonResult {
	essays, paging := services.TechService.FindRelated(params.NewQueryParams(c.Ctx).
		EqByReq("type").PageByReq().Desc("id"))
	return web.JsonPageData(essays, paging)
}

// GetRanked 文章排行榜
func (c *TechController) GetRanked() *web.JsonResult {
	essays, paging := services.TechService.FindRanked(params.NewQueryParams(c.Ctx).PageByReq().Desc("hits"))
	return web.JsonPageData(essays, paging)
}

// GetHit 点击量加1
func (c *TechController) GetHit() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonError(err)
	}
	e := services.TechService.AddHit(id)
	if e != nil {
		return web.JsonError(e)
	}
	return web.JsonSuccess()
}

// GetLike 喜爱次数加1
func (c *TechController) GetLike() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonError(err)
	}
	e := services.TechService.AddLike(id)
	if e != nil {
		return web.JsonError(e)
	}
	return web.JsonSuccess()
}
