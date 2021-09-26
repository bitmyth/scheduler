package channel

import (
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

var emailer = &EmailSender{}

func init() {
	emailer.Start()
}

type EmailSender struct {
	Ch chan *gomail.Message
}

func Send(m *gomail.Message) { emailer.Send(m) }
func (e *EmailSender) Send(m *gomail.Message) {
	e.Ch <- m
}

func (e *EmailSender) Start() {
	e.Ch = make(chan *gomail.Message)

	go func() {
		d := gomail.NewDialer("smtp.163.com", 465, "fishis", "")

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-e.Ch:
				if !ok {
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						log.Print(err)
						continue
					}
					open = true
				}
				if err := gomail.Send(s, m); err != nil {
					log.Print(err)
				}
			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						log.Print(err)
					}
					open = false
				}
			}
		}
	}()
}
