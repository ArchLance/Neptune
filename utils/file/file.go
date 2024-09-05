package file

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func IsFileExist(fileDir string) (bool, error) {
	// 使用os.Stat获取文件或目录的信息
	info, err := os.Stat(fileDir)
	log.Info(info)
	if err != nil {
		// 如果返回的错误是os.ErrNotExist，则表示文件或目录不存在
		if os.IsNotExist(err) {
			return false, nil // 文件或目录不存在，但这不是一个错误，所以我们返回nil作为错误
		}
		// 如果返回的是其他错误，则直接返回
		return false, err
	}
	// 如果文件或目录存在，则返回true和nil
	return info.IsDir(), nil // 注意：这里我们检查是否是目录，并返回其相反值。如果你只关心存在性而不区分文件或目录，可以只返回true
}
