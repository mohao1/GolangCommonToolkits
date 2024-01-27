package LinkList

type Node[T any] struct {
	val  T
	Next *Node[T]
}

type LinkList[T any] struct {
	node   *Node[T]
	length int
}

func (l *LinkList[T]) Get(index int) T {
	if l.IsEmpty() || index > l.length || index == 0 {
		return *new(T)
	}

	pre := l.node
	for i := 0; i < index; i++ {
		pre = pre.Next
	}

	return pre.val
}

func (l *LinkList[T]) Add(val T) int {
	pre := l.node
	for i := 0; i < l.length; i++ {
		pre = pre.Next
	}

	pre.Next = &Node[T]{
		val:  val,
		Next: nil,
	}
	l.length++
	return l.length
}

func (l *LinkList[T]) Remove(index int) T {
	if l.IsEmpty() || index > l.length || index == 0 {
		return *new(T)
	}

	pre := l.node

	for i := 0; i < index-1; i++ {
		pre = pre.Next
	}
	i := pre.Next
	pre.Next = i.Next
	l.length--

	return i.val
}

func (l *LinkList[T]) UpData(index int, val T) bool {
	if l.IsEmpty() || index > l.length || index == 0 {
		return false
	}
	pre := l.node
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	pre.val = val
	return true
}

func (l *LinkList[T]) Insert(index int, val T) bool {
	if index > l.length+1 {
		return false
	}
	pre := l.node
	for i := 0; i < index-1; i++ {
		pre = pre.Next
	}
	node := Node[T]{
		val:  val,
		Next: pre.Next,
	}
	pre.Next = &node
	l.length++
	return true
}

func (l *LinkList[T]) GetLength() int {
	return l.length
}

func (l *LinkList[T]) IsEmpty() bool {
	return l.length == 0
}

func (l *LinkList[T]) IsNoEmpty() bool {
	return !l.IsEmpty()
}

func NewLinkList[T any]() LinkList[T] {
	return LinkList[T]{
		length: 0,
		node: &Node[T]{
			val:  *new(T),
			Next: nil,
		},
	}
}
