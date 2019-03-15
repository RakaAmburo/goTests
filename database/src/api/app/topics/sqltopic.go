package topics

import (
	"github.com/mercadolibre/goTests/database/src/api/app"
	"sync"
)

type SqlTopic struct {
	packages chan app.Package
	list     []app.Package
	mutex    *sync.Mutex
}

func (t *SqlTopic) Init(size int) {
	t.packages = make(chan app.Package, size)
	t.list = make([]app.Package, size)
	t.mutex = &sync.Mutex{}
	go listen(t.packages, t.list, t.mutex)
}

func (t *SqlTopic) Publish(pkg app.Package) {
	t.packages <- pkg
}

func (t *SqlTopic) Get(index int) app.Package {
	t.mutex.Lock()
	result := t.list[index]
	t.mutex.Unlock()

	return result
}

func listen(packages <-chan app.Package, list []app.Package, mutex *sync.Mutex) {
	for pkg := range packages {
		index := pkg.GetIndex()
		mutex.Lock()
		list[index] = pkg
		mutex.Unlock()
	}
}
