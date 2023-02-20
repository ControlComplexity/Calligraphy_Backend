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
NewsController 书法教育文章，以图片为主
*/
type NewsController struct {
	Ctx iris.Context
}

// GetNews 文章列表
func (c *NewsController) GetNews() *web.JsonResult {
	newss, paging := services.NewsService.Find(params.NewQueryParams(c.Ctx).
		EqByReq("contentdynasty").EqByReq("writerdynasty").EqByReq("calligrapher").
		EqByReq("type").LikeByReq("title").PageByReq().Desc("id"))
	res := make([]model.News, 0)
	for _, es := range newss{
		res = append(res, model.News{
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
func (c *NewsController) GetBy(id int64) *web.JsonResult {
	newss := services.NewsService.Get(id)
	return web.JsonPageData(newss, nil)
}

// GetRelated 相关文章列表
func (c *NewsController) GetRelated() *web.JsonResult {
	newss, paging := services.NewsService.FindRelated(params.NewQueryParams(c.Ctx).
		EqByReq("type").PageByReq().Desc("id"))
	return web.JsonPageData(newss, paging)
}

// GetRanked 文章排行榜
func (c *NewsController) GetRanked() *web.JsonResult {
	newss, paging := services.NewsService.FindRanked(params.NewQueryParams(c.Ctx).PageByReq().Desc("hits"))
	return web.JsonPageData(newss, paging)
}

// GetHit 点击量加1
func (c *NewsController) GetHit() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonError(err)
	}
	e := services.NewsService.AddHit(id)
	if e != nil {
		return web.JsonError(e)
	}
	return web.JsonSuccess()
}

// GetLike 喜爱次数加1
func (c *NewsController) GetLike() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonError(err)
	}
	e := services.NewsService.AddLike(id)
	if e != nil {
		return web.JsonError(e)
	}
	return web.JsonSuccess()
}
