package main

import (
	"os"
	"time"
	"timeweel-job/job"
	"timeweel-job/scheduler"
)

type WriteMessage struct {
}

func (r *WriteMessage) Run() error {
	date := time.Now().Format("2016-01-02")

	filename := "msg" + date + ".txt"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	text := []byte(time.Now().String() + "\n")
	_, err = f.Write(text)
	return err
}

//func init() {
//_ = os.Setenv("TZ", "Asia/Shanghai")
//}

func main() {
	const TimeLayout = "2006-01-02 15:04:05"
	println(time.Now().String())

	timeStr := "2021-09-26 18:56:00"
	locationName := "Asia/Shanghai"
	l, _ := time.LoadLocation(locationName)
	//date, _ := time.Parse("2006-01-02 15:04:05", "2021-09-26 16:39:00")
	lt, _ := time.ParseInLocation(TimeLayout, timeStr, l)
	println(lt.String())


	j := &job.Job{
		RunAt: lt,
		Runnable: &job.Email{
			Text:    "timeweel job",
			Subject: "timeweel job",
		},
	}

	j.RepeatEveryNDays(1)
	//job.RepeatEveryWeek()
	//job.RepeatEveryMonth(5)

	timeWheel := scheduler.NewMinuteWheel()


	j2 := &job.Job{
		RunAt:    lt,
		Runnable: &WriteMessage{},
	}
	j.RepeatEveryNDays(1)

	timeWheel.Add(j)
	timeWheel.Add(j2)
	timeWheel.Start()

	//m := gomail.NewMessage()
	//m.SetHeader("From", "fishis@163.com")
	//m.SetHeader("To", "bitmyth@outlook.com", "fishis@163.com")
	//m.SetHeader("Subject", "Hello!")
	//m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//
	//d := gomail.NewDialer("smtp.163.com", 465, "fishis", "VDTCFGAPTPGLYFSW")
	//
	//// Send the email to Bob, Cora and Dan.
	//if err := d.DialAndSend(m); err != nil {
	//	panic(err)
	//}
}
