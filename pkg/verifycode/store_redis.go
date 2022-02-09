package verifycode

import (
	"go-web/pkg/app"
	"go-web/pkg/config"
	"go-web/pkg/redis"
	"time"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 插入值
func (redis *RedisStore) Set(id string, value string) bool {

	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	// 本地环境方便调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}

	return redis.RedisClient.Set(redis.KeyPrefix+id, value, ExpireTime)
}

// Get 获取缓存的值
func (redis RedisStore) Get(id string, clear bool) string {
	key := redis.KeyPrefix + id
	val := redis.RedisClient.Get(key)
	if clear {
		redis.RedisClient.Del(key)
	}
	return val
}

// Verify 验证是否存在
func (redis RedisStore) Verify(id, answer string, clear bool) bool {
	value := redis.Get(id, clear)
	return value == answer
}
