package tools

import (
	"fmt"
	"testing"
)

func Test_RandomWait(t *testing.T) {
	randWait := &RandomWait{}
	randWait.Init(100, 200)

	for i := 0; i < 20 ; i++ {
		fmt.Println(randWait.ShowWaitTime())
	}
}