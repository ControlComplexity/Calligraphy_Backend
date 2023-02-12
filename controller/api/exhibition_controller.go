package api

import (
	services "github.com/ControlComplexity/Calligraphy_Backend/service"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web"
)

type ExhibitionController struct {
	Ctx iris.Context
}

// GetExhibitions 文章列表
func (c *ExhibitionController) GetExhibitions() *web.JsonResult {
	essays, paging := services.ExhibitionService.Find(sqls.NewCnd().Desc("id"))
	return web.JsonPageData(essays, paging)
}

// GetBy 根据ID获取文章
func (c *ExhibitionController) GetBy(id int64) *web.JsonResult {
	essays := services.EssayService.Get(id)
	return web.JsonPageData(essays, nil)
}
