package redis

import "time"

func FindUser(key string) (int, error) {
	return RD.Get(key).Int()
}
func AddNotExistUser(key string) error {
	return RD.Set(key, "1", time.Minute).Err()
}
func FindRelation(key, member string) (bool, error) {
	return RD.SIsMember(key, member).Result()
}
func AddRelation(key, member string) (int64, error) {
	return RD.SAdd(key, member).Result()
}
func DelRelation(key, member string) (int64, error) {
	return RD.SRem(key, member).Result()
}
