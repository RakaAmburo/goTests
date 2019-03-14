package app

import (
	"sync"
)

type SqlTopic struct {
	packages chan Package
	list     []Package
	mutex    *sync.Mutex
}

func (t *SqlTopic) Init(size int) {
	t.packages = make(chan Package, size)
	t.list = make([]Package, size)
	t.mutex = &sync.Mutex{}
	go listen(t.packages, t.list, t.mutex)
}

func (t *SqlTopic) Publish(pkg Package) {
	t.packages <- pkg
}

func (t *SqlTopic) Get(index int) Package {
	t.mutex.Lock()
	result := t.list[index]
	t.mutex.Unlock()

	return result
}

func listen(packages <-chan Package, list []Package, mutex *sync.Mutex) {
	for pkg := range packages {
		index := pkg.GetIndex()
		mutex.Lock()
		list[index] = pkg
		mutex.Unlock()
	}
}
