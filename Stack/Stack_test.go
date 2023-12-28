package Stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	newStack := NewStack[string](10)
	newStack.Push("1")
	newStack.Push("3")
	newStack.Push("2")
	newStack.Push("6")
	for newStack.IsNoEmpty() {
		fmt.Println(newStack.Pop())
	}

}

func TestNewStack_Peek(t *testing.T) {
	newStack := NewStack[string](10)
	fmt.Println(newStack.Peek())
	newStack.Push("1")
	newStack.Push("3")
	fmt.Println(newStack.Peek())
	newStack.Push("2")
	newStack.Push("6")
	fmt.Println(newStack.Peek())
}
