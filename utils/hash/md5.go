package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5str(name string) string {
	res := md5.Sum([]byte(name))
	return hex.EncodeToString(res[:])
}
