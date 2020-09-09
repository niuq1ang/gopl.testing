package main

import (
	"fmt"
	"time"

	"gopkg.in/redis.v5"
)

type User struct {
	Name string
	Age  int
}

func main() {
	err := InitRedis()
	if err != nil {
		panic(err)
	}
	u := User{
		Name: "niuqiang",
		Age:  18,
	}
	RedisClient.Set("niuqiang_name", "u", time.Second*30)
	data, err := RedisClient.Get("niuqiang_name").Result()
	fmt.Println(data, err)
	u1 := new(User)
	RedisClient.HSet("niuqiang_name_1", "name", u)
	fmt.Println(RedisClient.HGet("niuqiang_name_1", "name").Scan(u1))
	fmt.Printf("%v", u1)

}

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return RedisClient.FlushDb().Err()
}
