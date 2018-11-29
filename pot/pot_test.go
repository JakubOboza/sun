package pot

import (
	"testing"
	"time"
)

type testTask struct {
	performed bool
	startAt   time.Time
}

func (tt *testTask) Perform() {
	tt.performed = true
}

func (tt *testTask) Interval() time.Duration {
	return time.Millisecond
}

func (tt *testTask) StartAt() time.Time {
	return tt.startAt
}

func TestPot(t *testing.T) {

	scheduler := Make()

	testTask := &testTask{performed: false, startAt: time.Now()}

	scheduler.AddTask(testTask)

	// pass time to sim normal usage
	time.Sleep(1 * time.Second)

	// run one internal iteration
	scheduler.next()

	if testTask.performed == false {
		t.Errorf("Task was not performed, got: %v, want: %v.", testTask.performed, true)
	}

}

func TestSimpleTaskNow(t *testing.T) {

	scheduler := Make()

	done := false

	testTask := MakeSimpleTaskNow(1, func() {
		done = true
	})

	scheduler.AddTask(testTask)

	// pass time to sim normal usage
	time.Sleep(1 * time.Second)

	// run one internal iteration
	scheduler.next()

	if done == false {
		t.Errorf("Task was not performed, got: %v, want: %v.", done, true)
	}

}
func TestSimpleTask(t *testing.T) {

	scheduler := Make()

	done := false

	testTask := MakeSimpleTask(time.Now().Add(5*time.Second), 1, func() {
		done = true
	})

	scheduler.AddTask(testTask)

	// pass time to sim normal usage
	time.Sleep(1 * time.Second)

	// run one internal iteration
	scheduler.next()

	if done == true {
		t.Errorf("Delayed start should be delayed but wasn't")
	}

	time.Sleep(5 * time.Second)

	scheduler.next()

	if done == false {
		t.Errorf("Delayed was not triggered after time passed")
	}

}
