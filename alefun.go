package alefun

import (
    "github.com/go-redis/redis"
)

type Alefun struct {
    redisClient *redis.Client
}

func NewAlefun (o *redis.Options) (*Alefun, error) {
    client := redis.NewClient(o)
    _, err := client.Ping().Result()
    if err != nil {
        return nil, err
    } else {
        return &Alefun{
            redisClient: client,
        }, nil
    }
}
