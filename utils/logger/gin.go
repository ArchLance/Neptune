package logger

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogFormatter(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	var output strings.Builder
	// 时间
	output.WriteString(param.TimeStamp.Format("[2006-01-02 15:04:05]"))
	// Code
	output.WriteString(statusColor)
	output.WriteRune('[')
	output.WriteString(strconv.Itoa(param.StatusCode))
	output.WriteRune(']')
	output.WriteString(resetColor)
	// Method
	output.WriteRune(' ')
	output.WriteString(methodColor)
	output.WriteString(param.Method)
	output.WriteString(resetColor)
	// Path
	output.WriteString(" - ")
	output.WriteString(param.Path)
	// 耗时
	output.WriteString(" | ")
	output.WriteString(param.Latency.String())
	// 错误
	output.WriteRune('\n')
	output.WriteString(param.ErrorMessage)
	return output.String()
}
