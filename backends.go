package alefun

import (
    "time"
    "github.com/go-redis/redis"
)

type Backend interface {
    Set(v interface{})
    SetExpires(v interface{}, expires time.Duration)
    Get(pk interface{}) interface{}
    Del(pk interface{})
}

type RedisBackend struct {
    redisClient *redis.Client
}

func (b *RedisBackend) Set(v interface{}) {

}
func (b *RedisBackend) SetExpires(v interface{}, expires time.Duration) {

}
func (b *RedisBackend) Get(pk interface{}) interface{} {
    return nil
}

func (b *RedisBackend) Del(pk interface{}) {

}
