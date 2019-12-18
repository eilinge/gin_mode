package main

import (
	"fmt"
	"time"
)

var timeLayoutStr = "2006-01-02 15:04:05.000"

func main02() {
	t1 := time.Now()
	time.Sleep(time.Second * 2)
	tt := time.Now().Sub(t1).Milliseconds()

	fmt.Println("the time: ", tt)
	fmt.Println("the time: float64", float64(tt))
}

func main05() {
	ts := TimeNowString()
	time.Sleep(1 * time.Second)
	fmt.Println("the value: ", getDuration(ts)) // 0: the duration cannot less 1 second
}

func main06() {
	ts := int64(time.Now().UnixNano() / (1000 * 1000)) // ms
	time.Sleep(2 * time.Millisecond)
	newSt := int64(time.Now().UnixNano() / (1000 * 1000)) // ms
	fmt.Println("the durateion :", newSt-ts)
}

// get duration of time.Now().Sub(p.Time)
func getDuration(ts string) int64 {
	st, _ := time.Parse(timeLayoutStr, ts)
	newSt, _ := time.Parse(timeLayoutStr, TimeNowString())
	fmt.Printf("the st: %v and the newSt: %v \n", st, newSt)
	return newSt.Sub(st).Milliseconds()
}

func TimeNowString() string {
	t := time.Now() // 当前时间
	t.UnixNano()
	fmt.Printf("the UTC value: %v \n", t)
	ts := t.Format(timeLayoutStr) // time转string
	fmt.Printf("the ts value: %v \n", ts)
	return t.String()
}
