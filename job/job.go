package job

import "time"

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
