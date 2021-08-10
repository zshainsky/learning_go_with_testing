package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

/* CONSTANTS */
const finalWord = "Go!"
const countDownStart = 3
const sleep = "sleep"
const write = "write"

/* INTERFACES */
type Sleeper interface {
	Sleep()
}

/* STRUCTS */
type SpySleeper struct {
	Calls int
}
type CountdownOperationsSpy struct {
	Calls []string
}
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
type SpyTime struct {
	durationSlept time.Duration
}

/* STRUCT IMPLEMENTATION */
func (s *SpySleeper) Sleep() {
	s.Calls++
}
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

/* APPLICATION FUNCTIONS */
func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

/* MAIN */
func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
