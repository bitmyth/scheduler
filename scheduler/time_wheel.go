package scheduler

import (
	"container/list"
	"fmt"
	"math"
	"time"
	"timeweel-job/job"
)

type TimeWheel struct {
	Slots        []*list.List
	currentSlot  int
	ticker       *time.Ticker
	tickDuration time.Duration
}

type Task struct {
	Job   *job.Job
	Round int
}

func (t *Task) SubtractRound() {
	if t.Round > 0 {
		t.Round--
	}
}

func New(slotCount int, tickDuration time.Duration) *TimeWheel {
	ticker := time.NewTicker(tickDuration)
	t := &TimeWheel{
		ticker:       ticker,
		tickDuration: tickDuration,
	}

	for i := 0; i < slotCount; i++ {
		t.Slots = append(t.Slots, list.New())
	}

	return t
}

func NewSecondWheel() *TimeWheel {
	slotCount := 3600
	return New(slotCount, time.Second)
}

func NewMinuteWheel() *TimeWheel {
	slotCount := 60
	return New(slotCount, time.Minute)
}

func (t TimeWheel) Start() {
	for tick := range t.ticker.C {
		t.currentSlot++
		t.currentSlot = t.currentSlot % len(t.Slots)
		fmt.Println("Tick at", tick, "current slot", t.currentSlot)

		node := t.Slots[t.currentSlot].Front()
		for node != nil {
			if node.Value != nil {
				task := node.Value.(*Task)

				fmt.Printf("%#v\n", task)
				if task.Round == 0 {
					fmt.Printf("Start Call %#v\n", task)
					go task.Job.Runnable.Run()

					if task.Job.IsRepeat {
						if task.Job.Period == job.Month {
							duration := time.Now().AddDate(0, 1, 0).Sub(time.Now())
							task.Round += int(duration/t.tickDuration) / len(t.Slots)
						} else {
							task.Round += int(task.Job.Period/t.tickDuration) / len(t.Slots)
						}
						println("set round", task.Round)
					} else {
						t.Slots[t.currentSlot].Remove(node)
					}
				}
				task.SubtractRound()
			}
			node = node.Next()
		}
	}
}

func (t *TimeWheel) Add(job *job.Job) {
	sub := job.RunAt.Sub(time.Now())
	var ticks int
	switch t.tickDuration {
	case time.Second:
		ticks = int(math.Ceil(sub.Seconds()))
	case time.Minute:
		ticks = int(math.Ceil(sub.Minutes()))
	}

	round := ticks / len(t.Slots)
	slotIndex := ticks % len(t.Slots)

	println("ticks", ticks, "round", round, "slotIndex", slotIndex)

	task := &Task{
		Job:   job,
		Round: round,
	}
	t.Slots[slotIndex].PushBack(task)
}
