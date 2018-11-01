package main

import (
	"fmt"
	"time"

	"github.com/JakubOboza/sun/pot"
)

// import (
// 	"fmt"
// 	"time"

// 	"github.com/JakubOboza/sun/relay_control"
// 	"github.com/JakubOboza/sun/vcgencmd"
// )

type PrintTask struct {
	startAt          time.Time
	body             string
	seconds_interval int
}

func (pt *PrintTask) Interval() time.Duration {
	return time.Duration(pt.seconds_interval) * time.Second
}

func (pt *PrintTask) StartAt() time.Time {
	return pt.startAt
}

func (pt *PrintTask) Perform() {
	fmt.Printf("%s %v\n", pt.body, time.Now())
}

func main() {

	// fmt.Printf("Hello Sun!\n")

	// relay_control.Open()

	// defer relay_control.Close()

	// pin := relay_control.New(22)

	// pin.TurnOn()
	// time.Sleep(6 * 10 * 100 * time.Millisecond)
	// pin.TurnOff()

	// float_temp, _ := vcgencmd.MeasureTemp()

	// fmt.Printf("( %v )\n", float_temp)

	scheduler := pot.Make()

	scheduler.AddTask(&PrintTask{body: "Yo!", seconds_interval: 1, startAt: time.Now()})
	scheduler.AddTask(&PrintTask{body: "Bro Staph!!!", seconds_interval: 6, startAt: time.Now()})
	scheduler.AddTask(&PrintTask{body: "So Booooring!", seconds_interval: 10, startAt: time.Now()})
	scheduler.AddTask(&PrintTask{body: "Fast one bro", seconds_interval: 2, startAt: time.Now()})

	scheduler.Run()
}
