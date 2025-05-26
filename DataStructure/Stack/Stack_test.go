package Stack

import (
	"common-toolkits-v1/DataStructure/Stack/LinkStack"
	"common-toolkits-v1/DataStructure/Stack/Stack"
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	newStack := Stack.NewStack[string](10)
	newStack.Push("1")
	newStack.Push("3")
	newStack.Push("2")
	newStack.Push("6")
	for newStack.IsNoEmpty() {
		fmt.Println(newStack.Pop())
	}

}

func TestNewStack_Peek(t *testing.T) {
	newStack := Stack.NewStack[string](10)
	fmt.Println(newStack.Peek())
	newStack.Push("1")
	newStack.Push("3")
	fmt.Println(newStack.Peek())
	newStack.Push("2")
	newStack.Push("6")
	fmt.Println(newStack.Peek())
}

func TestLinkNewStack_Peek(t *testing.T) {
	newlinkStack := LinkStack.NewLinkStack[string](4)
	fmt.Println("----")
	fmt.Println(newlinkStack.Peek())
	fmt.Println(newlinkStack.Push("1"))
	newlinkStack.Push("6")
	fmt.Println(newlinkStack.Peek())
	newlinkStack.Push("2")
	fmt.Println(newlinkStack.Pop())
	fmt.Println(newlinkStack.Pop())
	fmt.Println(newlinkStack.Pop())
	fmt.Println(newlinkStack.Pop())
}
