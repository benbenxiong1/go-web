package bootstrap

import (
	"fmt"
	"go-web/pkg/config"
	"go-web/pkg/redis"
)

func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"))
}
