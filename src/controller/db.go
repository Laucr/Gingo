package controller

import (
	"github.com/go-redis/redis"
	"fmt"
)

func Connect(db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: db})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("error", err)
		return nil
	}

	fmt.Println(pong)
	return client
}
