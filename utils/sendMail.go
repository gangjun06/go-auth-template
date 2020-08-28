package utils

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/mailgun/mailgun-go"
)

func SendVefiryMail(verifyCode, mail string) error {

	_mailHTML, errLoadMailHTML := ioutil.ReadFile("public/verifyEmail.html")
	if errLoadMailHTML != nil {
		log.Println(errLoadMailHTML)
	}

	config := GetConfig()

	mailHTML := string(_mailHTML)
	mailHTML = strings.Replace(mailHTML, "${link}", config.BaseURL+"/v1/auth/verify/"+verifyCode, -1)

	domain := config.MailDomain
	mg := mailgun.NewMailgun(domain, config.MailgunApiKey)
	sender := "NickName <no-reply@" + domain + ">"
	subject := "Title"
	body := ""

	message := mg.NewMessage(sender, subject, body, mail)
	message.SetHtml(mailHTML)

	_, _, errSendMail := mg.Send(message)
	return errSendMail
}

func SendPasswordReset(code, mail string) error {
	_mailHTML, errLoadMailHTML := ioutil.ReadFile("public/passwordCode.html")
	if errLoadMailHTML != nil {
		log.Println(errLoadMailHTML)
	}

	config := GetConfig()

	mailHTML := string(_mailHTML)
	mailHTML = strings.Replace(mailHTML, "${code}", code, -1)

	domain := config.MailDomain
	mg := mailgun.NewMailgun(domain, config.MailgunApiKey)
	sender := "NickName <no-reply@" + domain + ">"
	subject := "Title"
	body := ""

	message := mg.NewMessage(sender, subject, body, mail)
	message.SetHtml(mailHTML)

	_, _, errSendMail := mg.Send(message)
	return errSendMail
}
