package main

import (
	"flag"
	controllers "github.com/ControlComplexity/Calligraphy_Backend/controller"
	"github.com/ControlComplexity/Calligraphy_Backend/cron"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/ControlComplexity/Calligraphy_Backend/pkg/config"
	"github.com/mlogclub/simple/sqls"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"os"
	"time"
)

var configFile = flag.String("config", "./calligraphy.yml", "配置文件路径")

func init() {
	flag.Parse()

	// 初始化配置
	conf := config.Init(*configFile)

	// 初始化日志
	if file, err := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Error(err)
	}

	// 连接数据库
	gormConf := &gorm.Config{
		Logger: logger.New(logrus.StandardLogger(), logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
		}),
	}
	if err := sqls.Open(conf.DB.Url, gormConf, conf.DB.MaxIdleConns, conf.DB.MaxOpenConns, model.Models...); err != nil {
		logrus.Error(err)
	}
}
func main() {
	cron.RunCronJobs()
	controllers.Router()

}
