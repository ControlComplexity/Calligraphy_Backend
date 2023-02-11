package api

import (
	services "github.com/ControlComplexity/Calligraphy_Backend/service"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"github.com/mlogclub/simple/web/params"
)

type TheoryController struct {
	Ctx iris.Context
}

// GetTheory 轮播图列表
func (c *TheoryController) GetTheory() *web.JsonResult {
	essays, paging := services.TheoryService.Find(params.NewQueryParams(c.Ctx).
		EqByReq("type").LikeByReq("title").PageByReq().Desc("id"))
	return web.JsonPageData(essays, paging)
}
