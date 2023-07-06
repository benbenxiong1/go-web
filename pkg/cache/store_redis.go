package cache

import (
	"go-web/pkg/config"
	"go-web/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	return &RedisStore{
		RedisClient: redis.NewClient(address, username, password, db),
		KeyPrefix:   config.GetString("app.name") + ":cache:",
	}
}

func (s *RedisStore) Set(key string, value string, expireTime time.Duration) {
	s.RedisClient.Set(s.KeyPrefix+key, value, expireTime)
}

func (s *RedisStore) Get(key string) string {
	return s.RedisClient.Get(s.KeyPrefix + key)
}

func (s *RedisStore) Has(key string) bool {
	return s.RedisClient.Has(s.KeyPrefix + key)
}

func (s *RedisStore) Forget(key string) {
	s.RedisClient.Del(s.KeyPrefix + key)
}

func (s *RedisStore) Forever(key string, value string) {
	s.RedisClient.Set(s.KeyPrefix+key, value, 0)
}

func (s *RedisStore) Flush() {
	s.RedisClient.FlushDB()
}

func (s *RedisStore) IsAlive() error {
	return s.RedisClient.Ping()
}

// Increment 自增
func (s *RedisStore) Increment(parameters ...interface{}) {
	s.RedisClient.Incr(parameters...)
}

// Decrement 自减
func (s *RedisStore) Decrement(parameters ...interface{}) {
	s.RedisClient.Decr(parameters...)
}
