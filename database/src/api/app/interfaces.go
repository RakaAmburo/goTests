package app

type Writer interface {
	Write(line []string)
	Close()
	BulkWrite(pkg Package)
}

type Package interface {
	GetIndex() int
	Extract() [][]string
	Put(line []string)
}

type Topic interface {
	Publish(pkg Package)
	Get(index int) Package
}
