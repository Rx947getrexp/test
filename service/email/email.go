package email

import (
	"fmt"
	"go-speed/global"
	"net/smtp"
)

//func plainAuth() smtp.Auth {
//	identity := auth.identity
//	username := auth.username
//	password := auth.password
//	hostname := auth.hostname
//
//	return smtp.PlainAuth(identity, username, password, hostname)
//}

func loginAuth() smtp.Auth {
	username := auth.username
	password := auth.password

	return &LoginAuth{username, password}
}

func SendEmail(subject, body string, address []string) error {
	hostname := auth.hostname
	authentication := loginAuth()
	from := sender.addr
	//to := recipients.addr
	//msg := message.msg
	nickname := myEmailNickname
	contentType := "Content-Type: text/html; charset=UTF-8"
	for _, to := range address {
		s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s", to, nickname, from, subject, contentType, body)
		msg := []byte(s)
		err := smtp.SendMail(
			hostname,
			authentication,
			from,
			[]string{to},
			msg,
		)
		if err != nil {
			//log.Fatalln("Error!", err.Error())
			global.Logger.Err(err).Msg("发送邮件错误")
			return err
		}
	}

	fmt.Println("Awesome! Your email has been sent to the recipient.")
	return nil
}
