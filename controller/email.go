package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	"github.com/danangkonang/ceodeaja-go/helper"
)

// const CONFIG_SMTP_HOST = "smtp.gmail.com"
// const CONFIG_SMTP_PORT = 587
// const CONFIG_EMAIL = "ceodeaja@gmail.com"
// const CONFIG_PASSWORD = "rkdmeozhrwtkclso" //dafault
// // const CONFIG_PASSWORD = "vfuanendooviqpvo" //apps
const CONFIG_SMTP_HOST = "smtp.mailgun.org"
const CONFIG_SMTP_PORT = 587
const CONFIG_EMAIL = "ceodeaja@sandbox43c24b49c84e404187125d4cf9149471.mailgun.org"
const CONFIG_PASSWORD = "997ea96261648ee69bc47b01ddc925ca-7fba8a4e-0cccbdb5" //dafault

func SendEmail(w http.ResponseWriter, r *http.Request) {
	otp := helper.MakeOtp(6)
	to := []string{"dngrifai@gmail.com"}
	subject := "Test mail"
	message := HtmlCreate(otp)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// message := mime + "\r\n<html><body><h1>Hello World!" + otp + "</h1></body></html>"

	// message := "Hai user \n" +
	// 	"Ini adalah Kode otp ceodeaja: " + otp + "\n" +
	// 	"jangam memberi tahu siapa saja"
	err := sendMail(to, subject, mime, message)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Mail sent!")
}

func sendMail(to []string, subject, mime, message string) error {
	// body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body

	body := "From: " + CONFIG_EMAIL + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		mime + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_EMAIL, CONFIG_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)
	err := smtp.SendMail(smtpAddr, auth, CONFIG_EMAIL, append(to), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func HtmlCreate(otp string) string {
	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{.Title}}</title>
		</head>
		<body>
		<b>tes</b><br/>
			{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
		</body>
	</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
			otp,
		},
	}

	err = t.Execute(os.Stdout, data)
	check(err)
	return tpl
	// noItems := struct {
	// 	Title string
	// 	Items []string
	// }{
	// 	Title: "My another page",
	// 	Items: []string{},
	// }

	// err = t.Execute(os.Stdout, noItems)
	// check(err)
}
