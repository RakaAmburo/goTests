package app

import (
	"sync/atomic"
)

type Runner interface {
	run()
}

type Workers struct {
	jobs chan Runner
	executedTasks *uint64

}

func (w *Workers) Init(size int, workers int, timeOut *RandomWait) {
	w.jobs = make(chan Runner, size)
	w.executedTasks = new(uint64)
	for x := 1; x <= workers; x++ {
		go startWorkers(w.executedTasks, w.jobs, timeOut)
	}
	//close(w.jobs)
}

func startWorkers(counter *uint64, jobs <-chan Runner, timeOut *RandomWait) {
	for j := range jobs {
		j.run()
		atomic.AddUint64(counter, 1)
		timeOut.Wait()
	}
}

func (w *Workers) AddWork(job Runner){
	w.jobs <- job
}

func (w *Workers) GetExecutedTaskSize() uint64{
	return atomic.LoadUint64(w.executedTasks)
}
