package email

import (
	"strconv"
)

var (
	//env = getEnv()

	//myEmailAddr               = "ispeedbar@outlook.com"
	//myEmailPwrd               = "ispeed@bar"
	//myEmailNickname           = "一十九"
	myEmailAddr               = "VPNHERO@outlook.com"
	myEmailPwrd               = "pingguoqm23"
	myEmailNickname           = "HERO VPN"
	recipientEmailAddr string = "aaa@qq.com"

	sender = Sender{
		addr: myEmailAddr,
	}

	recipients = Recipients{
		addr: []string{recipientEmailAddr},
	}
)

var (
	subject string = "Testing function"
	body    string = "This is the email body."

	message = Message{
		msg: []byte("From: " + sender.addr + "\r\n" +
			"To: " + recipients.addr[0] + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n"),
	}
)

var (
	server = Server{
		name: "smtp-mail.outlook.com",
		port: 587,
	}

	auth = Auth{
		identity: "",
		username: myEmailAddr,
		password: myEmailPwrd,
		hostname: server.name + ":" + strconv.Itoa(server.port),
	}
)
