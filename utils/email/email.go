package email

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"neptune/global"
	"net/smtp"
)

func SendEmail(to string, subject string, content string) error {
	auth := smtp.PlainAuth("", global.ServerConfig.MailConfig.User, global.ServerConfig.MailConfig.AuthCode, global.ServerConfig.MailConfig.Host)
	address := fmt.Sprintf("%s:%s", global.ServerConfig.MailConfig.Host, global.ServerConfig.MailConfig.Port)
	e := email.NewEmail()
	//设置发送方的邮箱
	from := fmt.Sprintf("admin <%s>", global.ServerConfig.MailConfig.User)
	e.From = from
	// 设置接收方的邮箱
	e.To = []string{to}
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	e.Text = []byte(content)
	//err := global.EmailPool.Send(e, 120*time.Second)
	err := e.SendWithTLS(address, auth, &tls.Config{ServerName: global.ServerConfig.MailConfig.Host})
	return err
}
