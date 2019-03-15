package tools

import (
	"math/rand"
	"time"
)

type RandomWait struct {
	min int
	max int
}

func (rw *RandomWait) Init(min int, max int) {
	rw.min = min
	rw.max = max
}

func (rw *RandomWait) Wait() {
	waitTime := rand.Intn(rw.max-rw.min) + rw.min
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
}

func (rw *RandomWait) ShowWaitTime() int {
	rand.Seed(time.Now().UnixNano())
	waitTime := rand.Intn(rw.max-rw.min) + rw.min
	return waitTime
}
