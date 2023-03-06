package common

import (
	"errors"
	"github.com/go-redis/redis"
)

var (
	RepeatedFollow = errors.New("不能重复进行该操作")
	RedisNil       = redis.Nil
)
