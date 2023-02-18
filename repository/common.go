package repositories

import "github.com/go-redis/redis/v8"

var client *redis.Client
func getRedis() *redis.Client{
	if client == nil{
		client = redis.NewClient(&redis.Options{
			Addr:    "101.43.39.188:6379",   // 对应的ip以及端口号
			DB:       0,				 // 数据库的编号，默认的话是0
		})
	}
    return client
}
