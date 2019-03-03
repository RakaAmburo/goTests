package app

import (
	"fmt"
	"sync/atomic"
)

type Runner interface {
	run()
}

type Job struct {
}

func (Job) run() {
	fmt.Println("DOING SOME WORK")
}

type Workers struct {
	jobs chan Runner
	executedTasks *uint64
}

func (w *Workers) Init(size int) {
	w.jobs = make(chan Runner, size)
	w.executedTasks = new(uint64)
	for x := 1; x <= size; x++ {
		go startWorkers(w.executedTasks, w.jobs) //, results)
	}
	//close(w.jobs)
}

func startWorkers(counter *uint64, jobs <-chan Runner) { //, results chan<- Runner) {
	for j := range jobs {
		j.run()
		atomic.AddUint64(counter, 1)
	}
}

func (w *Workers) AddWork(job Runner){
	w.jobs <- job
}

func (w *Workers) GetExecutedTaskSize() uint64{
	return atomic.LoadUint64(w.executedTasks)
}
