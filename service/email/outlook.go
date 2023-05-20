package email

import (
	"fmt"
	"net/smtp"
)

const (
	// 邮件服务器地址
	SMTP_MAIL_HOST = "smtp.qq.com"
	// 端口
	SMTP_MAIL_PORT = "587"
	// 发送邮件用户账号
	SMTP_MAIL_USER = "aaa"
	// 授权密码
	SMTP_MAIL_PWD = "aaa"
	// 发送邮件昵称
	SMTP_MAIL_NICKNAME = "testMY"
)

func SendMail(address []string, subject, body string) (err error) {
	// 认证, content-type设置
	auth := smtp.PlainAuth("", SMTP_MAIL_USER, SMTP_MAIL_PWD, SMTP_MAIL_HOST)
	contentType := "Content-Type: text/html; charset=UTF-8"

	// 因为收件人，即address有多个，所以需要遍历,挨个发送
	for _, to := range address {
		s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s", to, SMTP_MAIL_NICKNAME, SMTP_MAIL_USER, subject, contentType, body)
		msg := []byte(s)
		addr := fmt.Sprintf("%s:%s", SMTP_MAIL_HOST, SMTP_MAIL_PORT)
		err = smtp.SendMail(addr, auth, SMTP_MAIL_USER, []string{to}, msg)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return err
}
