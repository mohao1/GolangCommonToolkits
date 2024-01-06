package Queue

import (
	"common-toolkits-v1/Queue/Queue"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue.NewQueue[int](5)
	queue.Push(10)
	queue.Push(20)
	queue.Push(30)
	fmt.Println(queue.Back())
	queue.Push(40)
	queue.Push(50)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Back())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())

}
