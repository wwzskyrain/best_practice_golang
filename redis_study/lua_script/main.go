package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/pratice/golang/redis_study"
)

func main() {
	rdb := redis_study.GetDefaultRedisClient()
	ctx := context.Background()
	SetAddAndExpire(ctx, rdb, "key", []interface{}{"1", "2", "3"}, 20)
	// incrBy.Run(ctx, rdb, []string{"my_counter"}, 2).Int()
}

func SetAddAndExpire(ctx context.Context, rdb *redis.Client, key string, members []interface{}, expireSecond int) int {
	var addAndExpire = redis.NewScript(`
		local key = KEYS[1]
		local expireSecond = ARGV[1]
		local member = ARGV[2]
		
		local value = redis.call("SADD", key, member)
		if not value then
		  value = 0
		end
		redis.call("EXPIRE", key, expireSecond)
		return value
	`)

	i, err := addAndExpire.Run(ctx, rdb, []string{key}, expireSecond, members[0]).Int()
	if err != nil {
		fmt.Printf("err => ", err)
		return -1
	}
	return i
}

var incrBy = redis.NewScript(`
local key = KEYS[1]
local change = ARGV[1]

local value = redis.call("GET", key)
if not value then
  value = 0
end

value = value + change
redis.call("SET", key, value)

return value
`)

// local expireSecond = ARGV[1]
// redis.call("EXPIRE", key, expireSecond)
