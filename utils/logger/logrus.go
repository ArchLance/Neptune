package logger

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// GetWriter 获取日志输出writer
func GetWriter() io.Writer {
	writerLock.Lock()
	defer writerLock.Unlock()
	if writer == nil {
		logDir := "log"
		logDate := 15
		logf, err := rotatelogs.New(
			filepath.Join(logDir, "neptune-%Y-%m-%d.log"),
			rotatelogs.WithMaxAge(time.Duration(logDate)*24*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			log.Error("Get rotate logs file err: ", err)
			return os.Stdout
		}
		writer = io.MultiWriter(os.Stdout, logf)
	}
	return writer
}

var writer io.Writer
var writerLock sync.Mutex

var FlagLToLevel = map[string]log.Level{
	"debug":   log.DebugLevel,
	"info":    log.InfoLevel,
	"warn":    log.WarnLevel,
	"warning": log.WarnLevel,
	"error":   log.ErrorLevel,
}

// logrus 日志格式化

type SimpleFormatter struct{}

func (f SimpleFormatter) Format(entry *log.Entry) ([]byte, error) {
	var output bytes.Buffer
	// 时间
	output.WriteString(entry.Time.Format("[2006-01-02 15:04:05]"))
	// 等级
	output.WriteRune('[')
	output.WriteString(entry.Level.String())
	output.WriteRune(']')
	// 消息
	output.WriteString(": ")
	output.WriteString(entry.Message)
	// 键值对
	if len(entry.Data) > 0 {
		output.WriteString(" | ")
	}
	for k, val := range entry.Data {
		output.WriteString(k)
		output.WriteRune('=')
		output.WriteString(cast.ToString(val))
		output.WriteRune(',')
	}
	output.WriteRune('\n')
	return output.Bytes(), nil
}
