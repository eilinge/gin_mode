package main

import (
	"testing"
	"time"
)

type smsPacket struct {
	Seq  int           `json:"seq"`
	Time time.Duration `json:"duration"`
	ttl  time.Duration `json:"-"`
}

func main04() {

}

func duration(ts string) int64 {
	st, _ := time.Parse(timeLayoutStr, ts)
	newSt, _ := time.Parse(timeLayoutStr, TimeNowString())
	return newSt.Sub(st).Milliseconds()
}

func TimeNowDuration() string {
	t := time.Now() // 当前时间
	t.Unix()

	ts := t.Format(timeLayoutStr) // time转string
	return ts
}

func Test_Millisecond(t *testing.T) {
	var waitFiveHundredMillisections time.Duration = 500 * time.Millisecond

	startingTime := time.Now().UTC()
	t.Logf("the value: %v", startingTime)
	time.Sleep(600 * time.Millisecond)
	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)

	if duration >= waitFiveHundredMillisections {
		t.Logf("Wait %v\nNative [%v]\nMilliseconds [%d]\nSeconds [%.3f]\n", waitFiveHundredMillisections, duration, duration.Nanoseconds()/1e6, duration.Seconds())
	}
}
