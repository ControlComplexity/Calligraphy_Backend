package cron

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/constant"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/go-redis/redis/v8"
	cron "github.com/robfig/cron/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"go.uber.org/zap"
)

var log *zap.Logger

var (
	client *redis.Client
	db       *gorm.DB
	linkOnce sync.Once
	redisOnce sync.Once
)
// 用于连接数据库
func linkDb() {
	linkOnce.Do(func() {
		dsn := "root:V1p3r1@#$%sdfvdswerDF43.@tcp(101.43.39.188:3306)/calligraphy?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		fmt.Println("redis Gorm")
		if err != nil {
			panic(err)
		}
	})
	redisOnce.Do(func() {
			client = redis.NewClient(&redis.Options{
				Addr:    "101.43.39.188:6379",   // 对应的ip以及端口号
				DB:       0,				 // 数据库的编号，默认的话是0
			})
			fmt.Println("redis NewClient")
	})
}

func init() {
	linkDb()
	log, _ = zap.NewProduction()
}
func addRank(members redis.Z) {
	// ZADD
	_, err := client.ZAdd( context.Background(), constant.CALLIGRAPHY_THEORY_ZSET_KEY, &members).Result()
	if err != nil {
		log.Error("failed to add data to Zset: "+ err.Error())
		return
	}
}

func addZset() {
	var essays []model.TheoryDO
	db.Select("id","title", "hits").Order("hits desc").Find(&essays)
	for _, item := range essays{
		bytes, err := json.Marshal(item)
		if err!= nil{
			log.Error("failed to marshal essay: "+ err.Error())
		}else{
			addRank(redis.Z{
				Member: string(bytes),
				Score: float64(item.Hits),
			})
		}
	}
	log.Info("finished to add Zset")
}

func RunCronJobs(){
	//启动的时候执行一次
	addZset()
	c := cron.New()
	//1. 每小时检查前五天和后五天的index是否存在，如果不存在的话创建
	c.AddFunc("@every 5m", func() {
		addZset()
	})
	c.Start()
}

