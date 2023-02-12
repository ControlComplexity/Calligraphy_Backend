package controllers

import (
	"github.com/ControlComplexity/Calligraphy_Backend/controller/api"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/mlogclub/simple/web"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
	ctx.Next()
}
func Router() {
	app := iris.New()
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(Cors)
	app.AllowMethods(iris.MethodOptions)

	app.OnAnyErrorCode(func(ctx iris.Context) {
		path := ctx.Path()
		var err error
		if strings.Contains(path, "/api/admin/") {
			err = ctx.JSON(web.JsonErrorCode(ctx.GetStatusCode(), "Http error"))
		}
		if err != nil {
			logrus.Error(err)
		}
	})

	app.Any("/", func(i iris.Context) {
		_, _ = i.HTML("<h1>Powered by Scenery</h1>")
	})

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/carousel").Handle(new(api.CarouselController))
		m.Party("/theory").Handle(new(api.TheoryController))
		m.Party("/essay").Handle(new(api.EssayController))
		m.Party("/activity").Handle(new(api.ActivityController))
		m.Party("/exhibition").Handle(new(api.ExhibitionController))
		m.Party("/tech").Handle(new(api.TechController))
	})

	if err := app.Listen(":10000",
		iris.WithConfiguration(iris.Configuration{
			DisableStartupLog:                 false,
			DisableInterruptHandler:           false,
			DisablePathCorrection:             false,
			EnablePathEscape:                  false,
			FireMethodNotAllowed:              false,
			DisableBodyConsumptionOnUnmarshal: false,
			DisableAutoFireStatusCode:         false,
			EnableOptimizations:               true,
			TimeFormat:                        "2006-01-02 15:04:05",
			Charset:                           "UTF-8",
		}),
	); err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}
