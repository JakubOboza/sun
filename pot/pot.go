package pot

import (
	"container/list"
	"time"
)

/*
  Pot is suppose to be a simple scheduler
*/

type Scheduler struct {
	tasks    *list.List
	commands chan Task
}

type taskWrapper struct {
	task      Task
	nextRunAt time.Time
}

func (tw *taskWrapper) reschedule() {
	tw.nextRunAt = time.Now().Add(tw.task.Interval())
}

func Make() *Scheduler {
	return &Scheduler{tasks: list.New(), commands: make(chan Task, 100)}
}

func (sc *Scheduler) AddTask(task Task) {
	taskItem := &taskWrapper{task: task}
	taskItem.nextRunAt = task.StartAt()
	sc.tasks.PushBack(taskItem)
}

func (sc *Scheduler) next() time.Duration {

	run_time := time.Now()
	min_time := run_time.Add(30 * time.Second) // max wait value to conserve the CPU

	for e := sc.tasks.Front(); e != nil; e = e.Next() {
		tw := e.Value.(*taskWrapper)
		if tw.nextRunAt.Before(run_time) {
			tw.task.Perform()
			//tw.nextRunAt = time.Now().Add(tw.task.Interval())
			tw.reschedule()
		}
		if min_time.After(tw.nextRunAt) {
			min_time = tw.nextRunAt
		}
	}
	// find all tasks to be run now
	// call perform on all of them
	// update the next run value!
	// calculate when next task should be run
	// return that duration value
	estimatedSleepTime := min_time.Sub(run_time)

	switch {
	case estimatedSleepTime > 60*time.Second:
		return 60 * time.Second
	case estimatedSleepTime < 4*time.Second:
		return 1 * time.Second
	default:
		return estimatedSleepTime
	}

}

func (sc *Scheduler) Run() {
	for {
		nextIterationIn := sc.next()
		select {
		case task := <-sc.commands:
			sc.AddTask(task)
		case <-time.After(nextIterationIn):
		}
	}
}
