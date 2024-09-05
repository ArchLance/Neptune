package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 两次哈希256后的字节数组，第二次是将第一次哈希后的16进制进行哈希
func SHA256Double(text string, isHex bool) []byte {
	hashInstance := sha256.New() // 实例赋值
	if isHex {
		arr, _ := hex.DecodeString(text) // 十六进制字符串转为十六进制字节数组
		hashInstance.Write(arr)          // 写入哈希实例对象
	} else {
		hashInstance.Write([]byte(text)) // 将字符串转换为字节数组，写入哈希对象
	}
	bytes := hashInstance.Sum(nil) // 哈希值追加到参数后面，只获取原始值，不用追加，用nil，返回哈希值字节数组
	hashInstance.Reset()           // 重置哈希实例
	hashInstance.Write(bytes)      // 将第一次哈希值写入哈希对象
	bytes = hashInstance.Sum(nil)  // 获取第二次哈希字节数组
	return bytes
}

// 两次哈希256后的哈希字符串，第二次是将第一次哈希后的16进制进行哈希
func SHA256DoubleString(text string, isHex bool) string {
	return fmt.Sprintf("%x", SHA256Double(text, isHex))
}
