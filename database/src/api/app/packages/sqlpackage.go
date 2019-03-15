package packages

type SqlPackage struct {
	list  [][]string
	index int
}

func (pkg *SqlPackage) Init(size int, index int) {// pasar tope?
	pkg.index = index
}

func (pkg *SqlPackage) Extract() [][]string {

	return pkg.list
}

func (pkg *SqlPackage) GetIndex() int {
	return pkg.index
}

func (pkg *SqlPackage) Put(line []string) {
	pkg.list = append(pkg.list, line)
}
