package app


type Writer interface {
	Write(line []string)
	Close()
	BulkWrite(pkg Package)
}

type Package interface {
	GetIndex() int
	Extract() [][]string
}

type Topic interface {
	Get(index int) Package
}