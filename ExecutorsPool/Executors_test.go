package ExecutorsPool

import (
	"fmt"
	"testing"
)

func TestExecutors_Run(t *testing.T) {
	executors := NewExecutors(2)
	executors.StartWorkerPool()

	for i := 0; i < 10; i++ {
		executors.Run(&Test{})
		executors.Run(&Test2{})
	}
	for i := 0; i < 10; i++ {
		executors.Run(&Test2{})
	}
	executors.GetWg().Wait()
}

type Test struct {
}

func (t *Test) GoRun() {
	fmt.Println("goRun")
}

type Test2 struct {
}

func (t *Test2) GoRun() {
	fmt.Println("goRun2")
}
