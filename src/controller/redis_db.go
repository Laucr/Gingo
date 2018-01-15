package controller

import (
	"github.com/go-redis/redis"
	"fmt"
	"sync"
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

func RedisInsert(db int, key string, fields map[string]interface{}) (int, int) {
	insertStatus := 0
	cli, e := connectRedis(db)
	if cli == nil {
		fmt.Println("Error:", e)
		return ConnectErr, OperationFailed
	}
	// confirm if exists
	val, err := cli.HGetAll(key).Result()
	if err != nil {
		return GetFailed, OperationFailed
	}
	if len(val) != 0 {
		insertStatus = InsertKeyExists
		return insertStatus, OperationFailed
	}

	// insert data
	var lock sync.Mutex
	lock.Lock()
	_, err = cli.HMSet(key, fields).Result()
	if err != nil {
		return InsertFailed, OperationFailed
	}
	lock.Unlock()

	// close client
	err = cli.Close()
	if err != nil {
		return CloseErr, OperationFailed
	}
	return InsertSuccess, OperationSuccess
}

func RedisLookup(db int, key string) (*map[string]string, int) {

	cli, e := connectRedis(db)
	if cli == nil {
		fmt.Println("Error:", e)
		return nil, ConnectErr
	}

	val, err := cli.HGetAll(key).Result()
	if err != nil {
		return nil, GetFailed
	}
	if len(val) == 0 {
		return nil, GetKeyNotExist
	}

	err = cli.Close()
	if err != nil {
		return nil, OperationFailed
	}

	return &val, OperationSuccess
}
