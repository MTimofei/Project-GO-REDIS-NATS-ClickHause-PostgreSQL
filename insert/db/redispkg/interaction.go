package redispkg

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// поключение к redis
func ConectRedis() (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}

// отправка данных в redis. принемает json
func Set(rdb *redis.Client, jsonByte []byte) (err error) {
	err = rdb.Set(context.Background(), "get", jsonByte, 60*time.Second).Err()
	if err != nil {
		err = fmt.Errorf("ERROR redis set: %q", err)
		return err
	}
	return nil
}

// получение даных из redis. отдает срез json
func Get(rdb *redis.Client) (result []byte, err error) {
	result, err = rdb.Get(context.Background(), "get").Bytes()
	if err != nil {
		err = fmt.Errorf("ERROR redis get: %q", err)
		return nil, err
	}

	return result, nil
}
