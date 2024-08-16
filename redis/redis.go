package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	db *redis.Client
)

func init() {
	db = Connect()
}

/**
 * @description:  连接redis数据库
 * @return {*}
 */
func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "127.0.0.1", "6379"),
		DB:   1,
	})
	return rdb
}

func main() {
	ctx := context.Background()
	ret := db.Conn().SAdd(ctx, "numbers", 1, 2, 3)

	fmt.Println(ret, db.Get(ctx, "numbers"))
}
