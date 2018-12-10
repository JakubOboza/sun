package pot

import "time"

func MakeSimpleTaskNow(intervalInSeconds int, foo func()) *SimpleTask {
	return &SimpleTask{startAt: time.Now(), intervalInSeconds: intervalInSeconds, foo: foo, repeat: true}
}

func MakeSimpleTask(startAt time.Time, intervalInSeconds int, foo func()) *SimpleTask {
	return &SimpleTask{startAt: startAt, intervalInSeconds: intervalInSeconds, foo: foo, repeat: true}
}

func MakeSingleRunSimpleTask(startAt time.Time, intervalInSeconds int, foo func()) *SimpleTask {
	return &SimpleTask{startAt: startAt, intervalInSeconds: intervalInSeconds, foo: foo, repeat: false}
}

type SimpleTask struct {
	startAt           time.Time
	intervalInSeconds int
	foo               func()
	repeat            bool
}

func (st *SimpleTask) Interval() time.Duration {
	return time.Duration(st.intervalInSeconds) * time.Second
}

func (st *SimpleTask) StartAt() time.Time {
	return st.startAt
}

func (st *SimpleTask) Perform() {
	st.foo()
}

func (st *SimpleTask) Repeat() bool {
	return st.repeat
}
