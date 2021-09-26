package job

import (
	"testing"
	"time"
)

func TestEmail_Run(t *testing.T) {
	e := Email{
		Subject: "Test",
		Text:    "Hello <b>Bob</b> and <i>hhhhhh</i>!",
	}
	e.Run()
	time.Sleep(50 * time.Second)
}
