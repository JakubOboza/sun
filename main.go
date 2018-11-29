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

type FanControl struct {
	fan      *relay_control.Relay
	hotTemp  float64
	coolTemp float64
	startAt  time.Time
}

func (fc *FanControl) Perform() {
	currentCpuTemp, _ := vcgencmd.MeasureTemp()
	if currentCpuTemp > fc.hotTemp {
		fmt.Printf("Need to chill %v\n", currentCpuTemp)
		fc.fan.TurnOn()
	}
	if currentCpuTemp < fc.coolTemp {
		fmt.Printf("Everything is fine %v\n", currentCpuTemp)
		fc.fan.TurnOff()
	}
}

func (fc *FanControl) StartAt() time.Time {
	return fc.startAt
}

func (fc *FanControl) Interval() time.Duration {
	return 10 * time.Second
}

func main() {

	// fmt.Printf("Hello Sun!\n")

	relay_control.Open()
	defer relay_control.Close()
	pin := relay_control.New(22)

	// pin.TurnOn()
	// time.Sleep(6 * 10 * 100 * time.Millisecond)
	// pin.TurnOff()

	// float_temp, _ := vcgencmd.MeasureTemp()

	// fmt.Printf("( %v )\n", float_temp)

	scheduler := pot.Make()

	fanController := &FanControl{fan: pin,hotTemp: 45.0, coolTemp: 39.2, startAt: time.Now()}
	scheduler.AddTask(fanController)

	// task1 := pot.MakeSimpleTaskNow(10, func() {
	// 	fmt.Println("Yo! \t", time.Now())
	// })

	// task2 := pot.MakeSimpleTaskNow(20, func() {
	// 	fmt.Println("Soup Man!! \t", time.Now())
	// })

	// task3 := pot.MakeSimpleTask(time.Now().Add(30*time.Second), 5, func() {
	// 	fmt.Println("My start was delayed by 30 sec\t", time.Now())
	// })

	// scheduler.AddTask(task1)
	// scheduler.AddTask(task2)
	// scheduler.AddTask(task3)

	cpuTempTask := pot.MakeSimpleTaskNow(5, func() {
		currentCpuTemp, _ := vcgencmd.MeasureTemp()
		fmt.Println("Cpu At: ", currentCpuTemp)
	})

	scheduler.AddTask(cpuTempTask)

	//scheduler.AddTask(&PrintTask{body: "Yo!", seconds_interval: 1, startAt: time.Now()})
	//scheduler.AddTask(&PrintTask{body: "Bro Staph!!!", seconds_interval: 6, startAt: time.Now()})
	//scheduler.AddTask(&PrintTask{body: "So Booooring!", seconds_interval: 10, startAt: time.Now()})
	//scheduler.AddTask(&PrintTask{body: "Fast one bro", seconds_interval: 2, startAt: time.Now()})

	scheduler.Run()
}
