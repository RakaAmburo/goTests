package app

import "fmt"

type Test struct {
	packages chan Package
	list    []Package
}

func (c *Test) Init(size int) {
	c.packages = make(chan Package, size)
	c.list = make([]Package, size)


	go listener(c.packages, c.list)
}

func (c *Test) Publish(pkg Package) {
	c.packages <- pkg
}

func (c *Test) Get(index int) Package{
	return c.list[index]
}

func listener(packages <-chan Package, list []Package) {
	for pkg := range packages {
		fmt.Println("pasa por aqui")
		index := pkg.GetIndex()
		list[index] = pkg
	}
}