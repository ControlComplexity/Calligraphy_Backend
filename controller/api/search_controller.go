package api

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	services "github.com/ControlComplexity/Calligraphy_Backend/service"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"github.com/mlogclub/simple/web/params"
)

/*
SearchController 关键字查询
*/
type SearchController struct {
	Ctx iris.Context
}

// GetResults 获取查询结果
func (c *SearchController) GetResults() *web.JsonResult {
	essays, paging := services.SearchService.Find(params.NewQueryParams(c.Ctx).
		EqByReq("type").LikeByReq("title").PageByReq().Desc("id"))
	res := make([]model.Theory, 0)
	for _, es := range essays{
		res = append(res, model.Theory{

		})
		fmt.Println("es: ", es)
	}
	return web.JsonPageData(res, paging)
}

// GetTop 获取查询查询次数最多的关键字
func (c *SearchController) GetTop() *web.JsonResult {
	essays := services.SearchService.Top()
	return web.JsonPageData(essays, nil)
}