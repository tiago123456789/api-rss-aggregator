package scheduler

import (
	"time"
)

func Task(intervalInSeconds int64, handler func()) {
	ticker := time.NewTicker(time.Second * time.Duration(intervalInSeconds))
	for ; ; <-ticker.C {
		handler()
	}
}
