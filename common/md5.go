package common

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}
