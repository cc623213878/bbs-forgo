package db

import (
	"bbs-forgo/utils/autoconfig"
	"fmt"
	"github.com/go-redis/redis"
)

//保存redis连接
var (
	RedisDb *redis.Client
)

//创建到redis的连接
func RedisDBInit() error {
	redisConf := autoconfig.Config.Base.Redis
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", redisConf.Host, redisConf.Port),
		Password: redisConf.Password, // no password set
		DB:       redisConf.DataBase, // use default DB
	})

	_, err := RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
