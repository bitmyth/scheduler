package main

import (
	"io/ioutil"
	"os"
	"time"
	"timeweel-job/job"
	"timeweel-job/scheduler"
)

type WriteMessage struct {
}

func (r *WriteMessage) Run() error {
	date := time.Now().Format("2016-01-02")
	ioutil.WriteFile("msg"+date+".txt", []byte(time.Now().String()), os.ModePerm)
	return nil
}

func main() {
	date, _ := time.Parse("2006-01-02 15:04:05", "2021-09-26 15:55:00")

	job := &job.Job{
		RunAt:    date,
		Runnable: &WriteMessage{},
	}

	job.RepeatEveryNDays(1)
	//job.RepeatEveryWeek()
	//job.RepeatEveryMonth(5)

	timeWheel := scheduler.NewMinuteWheel()
	timeWheel.Add(job)
	timeWheel.Start()
}
