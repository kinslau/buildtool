package util

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx context.Context

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "121.4.126.50:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func Connect() {

	ctx = context.Background()
	if err := initClient(); err != nil {
		return
	}

	err := rdb.Set(ctx, "kinslau", "itisme", 0).Err()
	if err != nil {
		panic(err)
	}
	set("kinslau", "test")
	get("kinslau")

}

func set(key string, val string) {
	err := rdb.SetNX(ctx, key, val, 10*time.Second).Err()
	fmt.Println(err)
	if err == redis.Nil {
		panic(err)
	} else if err != nil {
		panic(err)
	}

}

func get(key string) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("redis value: " + val)
	}

}
