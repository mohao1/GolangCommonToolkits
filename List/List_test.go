package List

import (
	"common-toolkits-v1/List/BLinkList"
	"common-toolkits-v1/List/LinkList"
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {
	list := LinkList.NewLinkList[int]()
	list.Add(10)
	list.Add(12)
	list.Add(11)
	list.Add(14)
	lens := list.GetLength()

	for i := 1; i <= lens; i++ {
		fmt.Println(list.Get(i))
	}
	list.Remove(1)
	lens = list.GetLength()
	for i := 1; i <= lens; i++ {
		fmt.Println(list.Get(i))
	}

	list.UpData(2, 100)
	for i := 1; i <= lens; i++ {
		fmt.Println(list.Get(i))
	}

}

func TestBLinkList(t *testing.T) {
	list := BLinkList.NewBLinkList[int]()
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)

	for i := 1; i <= list.GetLength(); i++ {
		fmt.Println(list.Get(i))
	}

	list.Remove(4)
	for i := 1; i <= list.GetLength(); i++ {
		fmt.Println(list.Get(i))
	}

	list.UpData(2, 100)
	for i := 1; i <= list.GetLength(); i++ {
		fmt.Println(list.Get(i))
	}
}
