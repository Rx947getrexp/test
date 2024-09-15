package email

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/global"
	"golang.org/x/net/context"
	"net"
	"net/smtp"
	"strings"
	"time"
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

func SetSendAccount(userName, passwd, hostName string) {
	auth.username = userName
	auth.password = passwd
	auth.hostname = hostName
}

func SendEmail(ctx *gin.Context, subject, body string, address []string) error {
	hostname := auth.hostname
	authentication := loginAuth()
	//from := sender.addr
	from := auth.username
	//to := recipients.addr
	//msg := message.msg
	nickname := myEmailNickname
	contentType := "Content-Type: text/html; charset=UTF-8"
	// 设置超时时间
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	resultCh := make(chan error)
	go func() {
		for _, to := range address {
			global.MyLogger(ctx).Info().Msgf("email from %s send to: %s", from, to)
			s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s", to, nickname, from, subject, contentType, body)
			msg := []byte(s)
			err := smtp.SendMail(
				//err := SendMailTLS(
				hostname,
				authentication,
				from,
				[]string{to},
				msg,
			)
			if err != nil {
				fmt.Println("email send to failed", err, to)
				global.MyLogger(ctx).Err(err).Msgf("email from %s send to: %s failed", from, to)
				resultCh <- err
				return
			} else {
				fmt.Println("email send to success", to)
				global.MyLogger(ctx).Info().Msgf("email from %s send to: %s success", from, to)
			}
		}
		resultCh <- nil
	}()
	// 等待 smtp.SendMail 完成或超时
	select {
	case <-ctxTimeout.Done():
		// 超时处理
		err := fmt.Errorf("smtp.SendMail timeout")
		fmt.Println(err.Error())
		global.MyLogger(ctx).Err(err).Msgf("email send timeout")
		return err
	case err := <-resultCh:
		// smtp.SendMail 完成
		if err != nil {
			fmt.Println("Error sending mail:", err)
			global.MyLogger(ctx).Err(err).Msgf("resultCh return failed")
		} else {
			fmt.Println("Mail sent successfully")
			global.MyLogger(ctx).Info().Msgf("Mail sent successfully")
		}
		return err
	}
}

func SendEmailTLS(ctx *gin.Context, subject, body string, address []string) error {
	hostname := auth.hostname
	authentication := loginAuth()
	from := auth.username
	nickname := myEmailNickname
	contentType := "Content-Type: text/html; charset=UTF-8"
	// 设置超时时间
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	resultCh := make(chan error)
	go func() {
		for _, to := range address {
			global.MyLogger(ctx).Info().Msgf("email from %s send to: %s", from, to)
			s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s", to, nickname, from, subject, contentType, body)
			msg := []byte(s)
			err := SendMailTLS(
				hostname,
				authentication,
				from,
				[]string{to},
				msg,
			)
			if err != nil {
				fmt.Println("email send to failed", err, to)
				global.MyLogger(ctx).Err(err).Msgf("email from %s send to: %s failed", from, to)
				resultCh <- err
				return
			} else {
				fmt.Println("email send to success", to)
				global.MyLogger(ctx).Info().Msgf("email from %s send to: %s success", from, to)
			}
		}
		resultCh <- nil
	}()
	// 等待 smtp.SendMail 完成或超时
	select {
	case <-ctxTimeout.Done():
		// 超时处理
		err := fmt.Errorf("smtp.SendMail timeout")
		fmt.Println(err.Error())
		global.MyLogger(ctx).Err(err).Msgf("email send timeout")
		return err
	case err := <-resultCh:
		// smtp.SendMail 完成
		if err != nil {
			fmt.Println("Error sending mail:", err)
			global.MyLogger(ctx).Err(err).Msgf("resultCh return failed")
		} else {
			fmt.Println("Mail sent successfully")
			global.MyLogger(ctx).Info().Msgf("Mail sent successfully")
		}
		return err
	}
}

// SendMailTLS not use STARTTLS commond
func SendMailTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	tlsconfig := &tls.Config{InsecureSkipVerify: true, ServerName: host}
	if err = validateLine(from); err != nil {
		return err
	}
	for _, recp := range to {
		if err = validateLine(recp); err != nil {
			return err
		}
	}
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	defer conn.Close()
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Hello("localhost"); err != nil {
		return err
	}
	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

// validateLine checks to see if a line has CR or LF as per RFC 5321
func validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return fmt.Errorf("a line must not contain CR or LF")
	}
	return nil
}
