package BLinkList

type Node[T any] struct {
	val T
	cur *Node[T] //后指针
	pre *Node[T] //前指针
}

type BLinkList[T any] struct {
	lNode  *Node[T]
	rNode  *Node[T]
	length int
}

func (B *BLinkList[T]) Get(index int) T {
	if B.IsEmpty() || index > B.length || index == 0 {
		return *new(T)
	}
	var pre *Node[T]
	if index <= B.length/2 {
		pre = B.lNode
		for i := 0; i < index; i++ {
			pre = pre.cur
		}
	} else {
		pre = B.rNode
		for i := 0; i < (B.length - index + 1); i++ {
			pre = pre.pre
		}
	}
	return pre.val

}

func (B *BLinkList[T]) Add(val T) int {
	node := &Node[T]{
		val: val,
		cur: B.rNode,
		pre: B.rNode.pre,
	}
	B.rNode.pre.cur = node
	B.rNode.pre = node
	B.length++
	return B.length
}

func (B *BLinkList[T]) Remove(index int) T {
	if B.IsEmpty() || index > B.length || index == 0 {
		return *new(T)
	}
	var pre *Node[T]
	if index <= B.length/2 {
		pre = B.lNode
		for i := 0; i < index; i++ {
			pre = pre.cur
		}
	} else {
		pre = B.rNode
		for i := 0; i < (B.length - index + 1); i++ {
			pre = pre.pre
		}
	}

	pre.pre.cur = pre.cur
	pre.cur.pre = pre.pre
	B.length--
	return pre.val
}

func (B *BLinkList[T]) UpData(index int, val T) bool {
	if B.IsEmpty() || index > B.length || index == 0 {
		return false
	}
	var pre *Node[T]
	if index <= B.length/2 {
		pre = B.lNode
		for i := 0; i < index; i++ {
			pre = pre.cur
		}
	} else {
		pre = B.rNode
		for i := 0; i < (B.length - index + 1); i++ {
			pre = pre.pre
		}
	}
	pre.val = val

	return true
}

func (B *BLinkList[T]) Insert(index int, val T) bool {
	if index > B.length+1 {
		return false
	}
	pre := B.lNode
	for i := 0; i < index; i++ {
		pre = pre.cur
	}
	node := &Node[T]{
		val: val,
		cur: pre,
		pre: pre.pre,
	}
	pre.pre.cur = node
	pre.pre = node
	B.length++
	return true
}

func (B *BLinkList[T]) GetLength() int {
	return B.length
}

func (B *BLinkList[T]) IsEmpty() bool {
	return B.length == 0
}

func (B *BLinkList[T]) IsNoEmpty() bool {
	return !B.IsEmpty()
}

func NewBLinkList[T any]() *BLinkList[T] {
	rNode := &Node[T]{}
	lNode := &Node[T]{}
	lNode.cur = rNode
	rNode.pre = lNode
	return &BLinkList[T]{
		length: 0,
		rNode:  rNode,
		lNode:  lNode,
	}
}
