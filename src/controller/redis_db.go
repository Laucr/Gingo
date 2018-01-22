package controller

import (
	"github.com/go-redis/redis"
	"fmt"
	"sync"
	"time"
)

func connectRedis(db int) (*redis.Client, int) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       db})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("error", err)
		return nil, ConnectErr
	}

	return client, OperationSuccess
}

func RedisInsert(db int, key string, j string, exp time.Duration) (int, int) {
	cli, e := connectRedis(db)
	if cli == nil {
		fmt.Println("Error:", e)
		return ConnectErr, OperationFailed
	}

	var lock sync.Mutex
	lock.Lock()
	_, err := cli.Set(key, j, exp).Result()
	if err != nil {
		return InsertFailed, OperationFailed
	}
	lock.Unlock()

	err = cli.Close()
	if err != nil {
		return CloseErr, OperationFailed
	}
	return InsertSuccess, OperationSuccess
}

func RedisSelect(db int, key string) (*string, int) {
	cli, e := connectRedis(db)
	if cli == nil {
		fmt.Println("Error:", e)
		return nil, ConnectErr
	}

	val, err := cli.Get(key).Result()
	if err != nil {
		return nil, QueryFailed
		fmt.Println(err)
	}
	if len(val) == 0 {
		return nil, QueryFailed
	}

	return &val, OperationSuccess
}

func CheckExpire(db int, key string) (int, int) {
	ttl := 0
	cli, e := connectRedis(db)
	if cli == nil {
		fmt.Println("Error:", e)
		return ttl, ConnectErr
	}

	dur, err := cli.TTL(key).Result()
	if err != nil {
		return -2, OperationFailed
	}
	return int(dur.Seconds()), OperationSuccess
}

func ExpireKey(db int, key string) (bool, int) {
	cli, e := connectRedis(db)
	if cli == nil {
		fmt.Println("Error:", e)
		return false, ConnectErr
	}

	res, err := cli.Expire(key, 0).Result()
	if err != nil {
		fmt.Println("Error", err)
		return res, OperationFailed
	}
	return res, OperationSuccess
}
