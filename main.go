package main

import (
	"fmt"
	"time"

	"github.com/JakubOboza/sun/vcgencmd"

	"github.com/JakubOboza/sun/relay_control"

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

type FanControl struct {
	fan      *relay_control.Relay
	hotTemp  float64
	coolTemp float64
	startAt  time.Time
}

func (fc *FanControl) Perform() {
	currentCpuTemp, _ := vcgencmd.MeasureTemp()
	if currentCpuTemp > fc.hotTemp {
		fmt.Printf("Need to chill %v", currentCpuTemp)
		fc.fan.TurnOn()
	}
	if currentCpuTemp < fc.coolTemp {
		fmt.Printf("Everything is fine %v", currentCpuTemp)
		fc.fan.TurnOff()
	}
}

func (fc *FanControl) StartAt() time.Time {
	return fc.startAt
}

func (fc *FanControl) Interval() time.Duration {
	return 5 * time.Second
}

func main() {

	// fmt.Printf("Hello Sun!\n")

	//relay_control.Open()
	//defer relay_control.Close()
	//pin := relay_control.New(22)

	// pin.TurnOn()
	// time.Sleep(6 * 10 * 100 * time.Millisecond)
	// pin.TurnOff()

	// float_temp, _ := vcgencmd.MeasureTemp()

	// fmt.Printf("( %v )\n", float_temp)

	scheduler := pot.Make()

	//fanController := &FanControl{hotTemp: 43.0, coolTemp: 40.0, startAt: time.Now()}
	//scheduler.AddTask(fanController)

	scheduler.AddTask(&PrintTask{body: "Yo!", seconds_interval: 1, startAt: time.Now()})
	scheduler.AddTask(&PrintTask{body: "Bro Staph!!!", seconds_interval: 6, startAt: time.Now()})
	scheduler.AddTask(&PrintTask{body: "So Booooring!", seconds_interval: 10, startAt: time.Now()})
	scheduler.AddTask(&PrintTask{body: "Fast one bro", seconds_interval: 2, startAt: time.Now()})

	scheduler.Run()
}
