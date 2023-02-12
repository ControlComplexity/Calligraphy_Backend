package api

import (
	services "github.com/ControlComplexity/Calligraphy_Backend/service"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"github.com/mlogclub/simple/web/params"
)

type ActivityController struct {
	Ctx iris.Context
}

// GetActivities 活动列表
func (c *ActivityController) GetActivities() *web.JsonResult {
	essays, paging := services.ActivityService.Find(params.NewQueryParams(c.Ctx).
		EqByReq("contentdynasty").EqByReq("writerdynasty").EqByReq("calligrapher").
		EqByReq("type").LikeByReq("title").PageByReq().Desc("activity_id"))
	return web.JsonPageData(essays, paging)
}

// GetBy 根据ID获取文章
func (c *ActivityController) GetBy(id int64) *web.JsonResult {
	essays := services.ActivityService.Get(id)
	return web.JsonPageData(essays, nil)
}
