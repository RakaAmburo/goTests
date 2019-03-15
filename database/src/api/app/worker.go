package app

import (
	"github.com/mercadolibre/goTests/database/src/api/app/tools"
	"sync/atomic"
)

type Runner interface {
	Run()
}

type Workers struct {
	jobs chan Runner
	executedTasks *uint64

}

func (w *Workers) Init(size int, workers int, timeOut *tools.RandomWait) {
	w.jobs = make(chan Runner, size)
	w.executedTasks = new(uint64)
	for x := 1; x <= workers; x++ {
		go startWorkers(w.executedTasks, w.jobs, timeOut)
	}
	//close(w.jobs)
}

func startWorkers(counter *uint64, jobs <-chan Runner, timeOut *tools.RandomWait) {
	for j := range jobs {
		j.Run()
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
