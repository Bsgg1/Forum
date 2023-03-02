package redis

import (
	"forum/common"
	"github.com/go-redis/redis"
)

func Rank(key, member string) (int64, error) {
	rank := RD.ZRank(key, member)
	if rank == nil {
		return -1, common.RepeatedFollow
	}
	return rank.Val(), nil
}
func AddLike(key, member string, score float64) error {
	return RD.ZAdd(key, redis.Z{
		Score:  score,
		Member: member,
	}).Err()

}

func DelLike() {

}
