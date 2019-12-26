package main

import (
	// "time"

	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func main07() {
	r := redisInit()
	_, err := r.Get("delTaskIDkey").Result()
	if err != nil {
		logrus.WithError(err).Error("redis get delivery reporter error")
		return
	}
}

func redisInit() redis.UniversalClient {
	c, _ := LoadConf("conf.yaml")
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        c.Redis.Addrs,
		Password:     c.Redis.Password,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})
}
