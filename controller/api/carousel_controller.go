package api

import (
	services "github.com/ControlComplexity/Calligraphy_Backend/service"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web"
)

type CarouselController struct {
	Ctx iris.Context
}

// GetCarousels 轮播图列表
func (c *CarouselController) GetCarousels() *web.JsonResult {
	imgs := services.CarouselService.Find(sqls.NewCnd().Desc("id"))
	return web.JsonPageData(imgs, nil)
}
