package reporters

import (
	"fmt"
	"sync"
)

var pkgPerWorker = make(map[int]*[]int)
var mutex sync.Mutex

func ReportPackagesPerWorker(workerNumber int, pkgNumber int) {

	mutex.Lock()
	value, ok := pkgPerWorker[workerNumber]
	if !ok {
		valueAux := make([]int, 0)
		value = &valueAux
		pkgPerWorker[workerNumber] = value
	}

	*value = append(*value, pkgNumber)
	mutex.Unlock()
}

func PrintPkgPerWrk() {
	for k, v := range pkgPerWorker {
		fmt.Print("Worker: ")
		fmt.Print(k)
		fmt.Printf("%v", *v)
		fmt.Println()
	}
}
