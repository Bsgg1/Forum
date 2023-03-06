package redis

import (
	"fmt"
	"time"
)

func AddPostsList(key, value string) {
	fmt.Println(key, value)
	RD.Set(key, value, time.Second)
}
func FindPostsList(key string) (string, error) {
	return RD.Get(key).Result()
}
func DelPostsList(key string) (int64, error) {
	return RD.Del(key).Result()
}
