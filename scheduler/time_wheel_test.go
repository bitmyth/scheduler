package scheduler

import (
	"testing"
	"time"
	"timeweel-job/job"
)

func TestDelayJob(t *testing.T) {
	timeWheel := NewSecondWheel()

	job := &job.Job{
		RunAt: time.Now().Add(time.Second * 15),
		Runnable: &job.RunPrint{
			Text: "aaa",
		},
	}

	timeWheel.Add(job)
	timeWheel.Start()
	time.Sleep(time.Second * 16)
}

func TestRepeatJob(t *testing.T) {
	timeWheel := NewSecondWheel()

	job := &job.Job{
		RunAt:    time.Now().Add(time.Second * 15),
		IsRepeat: true,
		Period:   10 * time.Second,
		Runnable: &job.RunPrint{
			Text: "aaa",
		},
	}

	timeWheel.Add(job)
	timeWheel.Start()
	time.Sleep(time.Second * 66)
}

func TestRepeatJob2(t *testing.T) {
	timeWheel := NewMinuteWheel()

	job := &job.Job{
		RunAt:    time.Now().Add(time.Second * 2),
		IsRepeat: true,
		Period:   1 * time.Hour,
		Runnable: &job.RunPrint{
			Text: "aaa",
		},
	}

	timeWheel.Add(job)
	timeWheel.Start()
	time.Sleep(time.Minute * 66)
}

func TestRepeatJob3(t *testing.T) {
	timeWheel := NewMinuteWheel()

	job := &job.Job{
		RunAt:    time.Now().Add(time.Second * 2),
		IsRepeat: true,
		Period:   1 * time.Hour,
		Runnable: &job.RunPrint{
			Text: "aaa",
		},
	}

	timeWheel.Add(job)
	timeWheel.Start()
	time.Sleep(time.Minute * 66)
}

func TestRepeatJob4(t *testing.T) {
	duration := time.Now().AddDate(0, 1, 0).Sub(time.Now()).Hours() / 24
	t.Log(duration)

	timeWheel := NewMinuteWheel()

	job := &job.Job{
		RunAt: time.Now().Add(time.Second * 2),
		Runnable: &job.RunPrint{
			Text: "aaa",
		},
	}
	//job.RepeatEveryWeek()
	job.RepeatEveryMonth(5)

	timeWheel.Add(job)
	timeWheel.Start()
	time.Sleep(time.Minute * 66)
}
func TestRepeatJob5(t *testing.T) {
	date, _ := time.Parse("2006-01-02 15:04:05", "2021-09-26 15:50:00")
	t.Log(date)
	job := &job.Job{
		RunAt: date,
		Runnable: &job.RunPrint{
			Text: "aaa",
		},
	}
	//job.RepeatEveryWeek()
	job.RepeatEveryMonth(5)

	timeWheel := NewMinuteWheel()
	timeWheel.Add(job)
	timeWheel.Start()
}
