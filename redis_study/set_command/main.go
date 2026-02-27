package main

import (
	"context"
	"fmt"

	"github.com/pratice/golang/redis_study"
)

func main() {
	redis := redis_study.GetDefaultRedisClient()
	ctx := context.Background()
	setKey := "s_k"
	// err := redis.SAdd(ctx, setKey, 1, 2, 3, 4).Err()
	// err := redis.SAdd(ctx, setKey, []int{1, 2, 3, 4}).Err()
	err := redis.SAdd(ctx, setKey, []interface{}{1, 2, 3, 4}).Err()
	if err != nil {
		fmt.Printf("sadd err => %+v \n", err)
		return
	}
	members := redis.SMembers(ctx, setKey)
	result, err1 := members.Result()
	if err1 != nil {
		fmt.Printf("err = %+v \n", err1)
		return
	}
	fmt.Printf("result = %+v", result)

}
