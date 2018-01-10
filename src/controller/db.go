package controller

import (
	"github.com/go-redis/redis"
	"fmt"
	"sync"
)

// Database operation return code
const (
	OperationSuccess = 0
	OperationFailed  = -1
	CloseErr         = -11
	ConnectErr       = -10
	InsertFailed     = -100
	InsertSuccess    = 100
	InsertKeyExist   = -101
	GetFailed        = -200
)

// Databases
const (
	UserInfo = 0
)

func Connect(db int) (*redis.Client, int) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       db})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("error", err)
		return nil, ConnectErr
	}

	return client, OperationSuccess
}

func CloseClient(cli *redis.Client) error {
	err := cli.Close()
	return err
}

func Insert(cli *redis.Client, key string, fields map[string]interface{}) (int, int) {
	insertStatus := 0

	// confirm if exists
	val, err := cli.HGetAll(key).Result()
	if err != nil {
		return GetFailed, OperationFailed
	}
	if len(val) != 0 {
		insertStatus = InsertKeyExist
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
	err = CloseClient(cli)
	if err != nil {
		return CloseErr, OperationFailed
	}
	return InsertSuccess, OperationSuccess
}
