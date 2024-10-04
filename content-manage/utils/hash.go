package utils

import (
	"hash/fnv"
	"math/big"
)

func GenIdx(uuid string, mod int) int {
	hash := fnv.New64()
	_, _ = hash.Write([]byte(uuid))
	hashVal := hash.Sum64()
	bigHashNum := big.NewInt(int64(hashVal))
	return int(bigHashNum.Mod(bigHashNum, big.NewInt(int64(mod))).Int64())

}
