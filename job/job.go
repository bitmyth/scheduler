package job

import (
	"gopkg.in/gomail.v2"
	"time"
	"timeweel-job/channel"
)

const (
	Day   = 24 * time.Hour
	Week  = 7 * Day
	Month = -1
)

type Job struct {
	RunAt    time.Time
	IsRepeat bool
	Period   time.Duration
	Runnable Runnable
	Date     int
}

func (j *Job) RepeatEveryNDays(days int) {
	j.IsRepeat = true
	j.Period = time.Duration(days) * Day
}

func (j *Job) RepeatEveryWeek() {
	j.IsRepeat = true
	j.Period = Week
}

// RepeatEveryMonth 每月几号重复
func (j *Job) RepeatEveryMonth(date int) {
	j.IsRepeat = true
	j.Period = Month
	j.Date = date
}

type Runnable interface {
	Run() error
}

type RunPrint struct {
	Text string
}

func (r *RunPrint) Run() error {
	println("text is :", r.Text)
	return nil
}

type Email struct {
	Subject string
	Text    string
}

func (e Email) Run() error {
	m := gomail.NewMessage()
	m.SetHeader("From", "fishis@163.com")
	m.SetHeader("To", "bitmyth@outlook.com", "fishis@163.com")
	m.SetAddressHeader("Cc", "1091301899@qq.com", "qq")
	m.SetHeader("Subject", e.Subject)
	m.SetBody("text/html", e.Text)
	//m.Attach("/home/Alex/lolcat.jpg")
	channel.Send(m)
	return nil
}