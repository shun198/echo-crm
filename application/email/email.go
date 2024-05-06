package email

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

type UrlData struct {
	Url string
}

type emailTemplates struct {
	html      string
	plaintext string
}

func sendEmail(e *echo.Echo, templates emailTemplates, email string, subject string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_ADMIN"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", templates.html)
	m.AddAlternative("text/plain", templates.plaintext)

	host := os.Getenv("MAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	d := gomail.NewDialer(host, port, username, password)

	if err := d.DialAndSend(m); err != nil {
		e.Logger.Error(err)
		return errors.New("メールの送信に失敗しました")
	}
	return nil
}

func sendEmailSES(e *echo.Echo, templates emailTemplates, email string, subject string) error {
	charset := "UTF-8"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)
	if err != nil {
		return err
	}

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charset),
					Data:    aws.String(templates.html),
				},
				Text: &ses.Content{
					Charset: aws.String(charset),
					Data:    aws.String(templates.plaintext),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charset),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(os.Getenv("MAIL_ADMIN")),
	}
	result, err := svc.SendEmail(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				e.Logger.Error(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				e.Logger.Error(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				e.Logger.Error(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				e.Logger.Error(aerr.Error())
			}
		} else {
			e.Logger.Error(err.Error())
		}

		return err
	}

	e.Logger.Info("Email Sent to address: " + email)
	e.Logger.Info(result)
	return nil
}

func CreateEmail(e *echo.Echo, email string, url string, templateName string, subject string) error {
	wd, _ := os.Getwd()

	html, _ := template.New(fmt.Sprint(templateName, ".html")).ParseFiles(fmt.Sprint(wd, "/templates/", templateName, ".html"))
	plaintext, _ := template.New(fmt.Sprint(templateName, ".txt")).ParseFiles(fmt.Sprint(wd, "/templates/", templateName, ".txt"))

	var tpl bytes.Buffer
	if err := html.Execute(&tpl, UrlData{Url: url}); err != nil {
		e.Logger.Error(err)
		return err
	}
	htmlResult := tpl.String()
	if err := plaintext.Execute(&tpl, UrlData{Url: url}); err != nil {
		e.Logger.Error(err)
		return err
	}
	plaintextResult := tpl.String()

	if os.Getenv("DEBUG") == "" {
		return sendEmailSES(e, emailTemplates{html: htmlResult, plaintext: plaintextResult}, email, subject)
	} else {
		return sendEmail(e, emailTemplates{html: htmlResult, plaintext: plaintextResult}, email, subject)
	}
}
