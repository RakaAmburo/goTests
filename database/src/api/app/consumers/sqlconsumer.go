package consumers

import (
	"github.com/mercadolibre/goTests/database/src/api/app"
	"sync"
	"time"
)

type SqlConsumer struct {
	pkg     chan app.Package
	counter int
	wait    *sync.WaitGroup
}

func (c *SqlConsumer) Init(size int, topic app.Topic, writer app.Writer, wait *sync.WaitGroup) {
	c.pkg = make(chan app.Package, size)
	c.counter = size
	c.wait = wait
	go consume(topic, writer, c, c.wait)
}

func consume(topic app.Topic, writer app.Writer, c *SqlConsumer, wait *sync.WaitGroup) {

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
