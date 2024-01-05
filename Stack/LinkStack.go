package Stack

type Node[T any] struct {
	Data    T
	curNode *Node[T]
}

type LinkStack[T any] struct {
	data *Node[T]
	size int
	top  int
}

func (s *LinkStack[T]) Push(value T) bool {
	if s.top == s.size {
		return false
	} else {
		pre := s.data
		for pre.curNode != nil {
			pre = pre.curNode
		}
		pre.curNode = &Node[T]{
			Data:    value,
			curNode: nil,
		}
		s.top++
		return true
	}
}

func (s *LinkStack[T]) Pop() T {
	if s.IsEmpty() {
		return *new(T)
	} else {
		pre := s.data
		for pre.curNode.curNode != nil {
			pre = pre.curNode
		}
		data := pre.curNode.Data
		pre.curNode = nil
		s.top--
		return data
	}
}

func (s *LinkStack[T]) Size() int {
	return s.top
}

func (s *LinkStack[T]) Peek() T {
	if s.IsEmpty() {
		return *new(T)
	} else {
		pre := s.data
		for pre.curNode.curNode != nil {
			pre = pre.curNode
		}
		data := pre.curNode.Data
		return data
	}
}

func (s *LinkStack[T]) IsEmpty() bool {
	return s.top == 0
}

func (s *LinkStack[T]) IsNoEmpty() bool {
	return !s.IsEmpty()
}

// NewLinkStack 获得一个顺序表栈
func NewLinkStack[T any](size int) *LinkStack[T] {
	return &LinkStack[T]{
		data: &Node[T]{
			Data:    *new(T),
			curNode: nil,
		},
		top:  0,
		size: size,
	}
}
