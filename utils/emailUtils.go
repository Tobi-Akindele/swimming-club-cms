package utils

import (
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"strconv"
)

func serverConfig() *mail.SMTPServer {
	server := mail.NewSMTPClient()
	server.Host = GetEnv(SMTP_HOST, "")
	server.Port, _ = strconv.Atoi(GetEnv(SMTP_PORT, ""))
	server.Username = GetEnv(SMTP_USER, "")
	server.Password = GetEnv(SMTP_PASS, "")
	server.Encryption = mail.EncryptionTLS

	return server
}

func SendMail(subject string, to string, message string) {
	smtpClient, err := serverConfig().Connect()
	if err != nil {
		log.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(GetEnv(FROM_EMAIL, ""))
	email.AddTo(to)
	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, message)

	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	}
}

func ComposeAccountActivationEmail(firstName string, activationCode string) string {
	setPasswordLink := GetEnv(SET_PASSWORD_PAGE_LINK, "") + activationCode
	htmlBody := fmt.Sprintf(`<html>
		<head>
			<meta http-equip="Content-Type" content="text/html; charset=utf-8" />
			<title>Account Activation</title>
		</head>
		<body>
			<p><h1>Swimming Club CMS</h1></p>
			<p></p>
			<p>Hello %s,</p>

			<p>Thank you for joining the Swimming Club CMS. Kindly click the button below to set your password and activate your account.</p>

			<p style="text-align:center;"><a href="%s" target="_blank" style="text-align:center; padding:10px; background-color:#79a5d9;text-decoration:none; color:white;">Set Password</a></p>

			<p>Or follow the link below</p>
			<p><a href="%s" target="_blank">%s</a></p>
			
			<p>Thank you.</p>
			<p></p>
			<p>Regards,</p>
			<p>Swimming Club CMS Team</p>
			<p>S405 PC Lab, Mellor Building, Staffordshire University.</p>
		</body>
		</html>`, firstName, setPasswordLink, setPasswordLink, setPasswordLink)

	return htmlBody
}
