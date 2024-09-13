package email

import (
	"github.com/jordan-wright/email"
	"neptune/global"
	"time"
)

func SendEmail(to string, subject string, content string) error {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "dj <XXX@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{"XXX@qq.com"}
	//设置主题
	e.Subject = "这是主题"
	//设置文件发送的内容
	e.Text = []byte("www.topgoer.com是个不错的go语言中文文档")
	err := global.EmailPool.Send(e, 10*time.Second)
	return err
}
