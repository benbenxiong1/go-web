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

func (s *RedisStore) Set(id string, value string) bool {
	expireKey := "verifycode.expire_time"
	if app.IsLocal() {
		expireKey = "verifycode.debug_expire_time"
	}

	expireTime := time.Minute * time.Duration(config.GetInt64(expireKey))

	return s.RedisClient.Set(s.KeyPrefix+id, value, expireTime)
}

func (s *RedisStore) Get(id string, clear bool) string {
	key := s.KeyPrefix + id
	value := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return value
}

func (s *RedisStore) Verify(id string, value string, clear bool) bool {
	v := s.Get(id, clear)
	return v == value
}
