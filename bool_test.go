package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func redisForTest(t *testing.T) redis.UniversalClient {
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

func Test_getchannelINfo(t *testing.T) {
	redis := redisForTest(t)
	rs, _ := redis.Get(getRedisKey("7377309771234304")).Result()
	var ch Channel
	_ = json.Unmarshal([]byte(rs), &ch)
	if ch.WindowStats {
		t.Error("unmarshal the channel windowstats type error")
	}
	fmt.Printf("the channel info:%#v", ch)
}
