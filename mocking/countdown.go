package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

// Spies are a kind of mock which can record how a dependency is used.
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}
