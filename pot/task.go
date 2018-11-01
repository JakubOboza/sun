package pot

import "time"

type Task interface {
	Perform()
	Interval() time.Duration
	StartAt() time.Time
}
