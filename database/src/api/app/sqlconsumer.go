package app

import (
	"sync"
	"time"
)

type SqlConsumer struct {
	pkg     chan Package
	counter int
	wait    *sync.WaitGroup
}

func (c *SqlConsumer) Init(size int, topic Topic, writer Writer, wait *sync.WaitGroup) {
	c.pkg = make(chan Package, size)
	c.counter = size
	c.wait = wait
	go consume(topic, writer, c, c.wait)
}

func consume(topic Topic, writer Writer, c *SqlConsumer, wait *sync.WaitGroup) {

	index := 0
	for index < c.counter {

		pkg := topic.Get(index)
		if pkg == nil {
			time.Sleep(100 * time.Millisecond)
		} else {
			index ++
			writer.BulkWrite(pkg)
		}

	}
	wait.Done()
}
