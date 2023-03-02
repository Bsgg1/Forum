package common

import "errors"

var (
	RepeatedFollow = errors.New("不能重复关注")
)
