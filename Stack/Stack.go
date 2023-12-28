package Stack

type Interface[T any] interface {
	Push(value T) bool //添加数据
	Pop() T            //弹出数据
	Size() int         //查看个数
	Peek() T           //查看栈顶的第一个元素
	IsEmpty() bool     //是否栈空
	IsNoEmpty() bool   //是否栈不为空
}

type Stack[T any] struct {
	data []T
	size int
	top  int
}

func (s *Stack[T]) Push(value T) bool {
	if s.top == s.size {
		return false
	}
	s.data[s.top] = value
	s.top++
	return true
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		return *new(T)
	} else {
		s.top--
		return s.data[s.top]
	}
}

func (s *Stack[T]) Size() int {
	return s.top
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		return *new(T)
	} else {
		return s.data[s.top-1]
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack[T]) IsNoEmpty() bool {
	return !s.IsEmpty()
}

// NewStack 获得一个顺序表栈
func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, size),
		size: size,
		top:  0,
	}
}
