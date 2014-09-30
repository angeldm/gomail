package gomail

import (
	"bytes"
	"log"
	"net/smtp"
	"strconv"
	"text/template"
)

type EmailUser struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}

type SmtpTemplateData struct {
	From    string
	To      string
	Subject string
	Body    string
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}

Sincerely,
`

type GoMail struct {
	emailUser     *EmailUser
	auth          smtp.Auth
	context       *SmtpTemplateData
	t             *template.Template
	configuracion *Configuration
}

func New() *GoMail {
	Gomail := &GoMail{}
	Gomail.configuracion = readConfig()
	Gomail.emailUser = &EmailUser{Gomail.configuracion.Username, Gomail.configuracion.Password, Gomail.configuracion.EmailServer, Gomail.configuracion.Port}
	Gomail.auth = smtp.PlainAuth("",
		Gomail.emailUser.Username,
		Gomail.emailUser.Password,
		Gomail.emailUser.EmailServer,
	)
	return Gomail

}

func NewWithConfigPaht(path string) *GoMail {
	Gomail := &GoMail{}
	Gomail.configuracion = readConfigWithPath(path)
	Gomail.emailUser = &EmailUser{Gomail.configuracion.Username, Gomail.configuracion.Password, Gomail.configuracion.EmailServer, Gomail.configuracion.Port}
	Gomail.auth = smtp.PlainAuth("",
		Gomail.emailUser.Username,
		Gomail.emailUser.Password,
		Gomail.emailUser.EmailServer,
	)
	return Gomail

}

func (g GoMail) sendMail(subject, body string) {

	var err error
	var doc bytes.Buffer

	context := &SmtpTemplateData{
		g.configuracion.Username,
		g.configuracion.To,
		subject,
		body,
	}

	g.t = template.New("emailTemplate")
	g.t, err = g.t.Parse(emailTemplate)

	if err != nil {
		log.Print("error trying to parse mail template")
	}
	log.Print("init ok")

	err = g.t.Execute(&doc, context)
	if err != nil {
		log.Print("error trying to execute mail template")
	}

	err = smtp.SendMail(g.configuracion.EmailServer+":"+strconv.Itoa(g.configuracion.Port), // in our case, "smtp.google.com:587"
		g.auth,
		g.configuracion.Username,
		[]string{g.configuracion.To},
		doc.Bytes())
	if err != nil {
		log.Print("ERROR: attempting to send a mail ", err)
	}
}
