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
EssayController 书法教育文章，以图片为主
*/
type EssayController struct {
	Ctx iris.Context
}

// GetEssays 文章列表
func (c *EssayController) GetEssays() *web.JsonResult {
	essays, paging := services.EssayService.Find(params.NewQueryParams(c.Ctx).
		EqByReq("contentdynasty").EqByReq("writerdynasty").EqByReq("calligrapher").
		EqByReq("type").LikeByReq("title").PageByReq().Desc("id"))
	res := make([]model.Essay, 0)
	for _, es := range essays{
		res = append(res, model.Essay{
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
func (c *EssayController) GetBy(id int64) *web.JsonResult {
	essays := services.EssayService.Get(id)
	return web.JsonPageData(essays, nil)
}

// GetRelated 相关文章列表
func (c *EssayController) GetRelated() *web.JsonResult {
	essays, paging := services.EssayService.FindRelated(params.NewQueryParams(c.Ctx).
		EqByReq("type").PageByReq().Desc("id"))
	return web.JsonPageData(essays, paging)
}

// GetRanked 文章排行榜
func (c *EssayController) GetRanked() *web.JsonResult {
	essays, paging := services.EssayService.FindRanked(params.NewQueryParams(c.Ctx).PageByReq().Desc("hits"))
	return web.JsonPageData(essays, paging)
}

// GetHit 点击量加1
func (c *EssayController) GetHit() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonError(err)
	}
	e := services.EssayService.AddHit(id)
	if e != nil {
		return web.JsonError(e)
	}
	return web.JsonSuccess()
}

// GetLike 喜爱次数加1
func (c *EssayController) GetLike() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonError(err)
	}
	e := services.EssayService.AddLike(id)
	if e != nil {
		return web.JsonError(e)
	}
	return web.JsonSuccess()
}
