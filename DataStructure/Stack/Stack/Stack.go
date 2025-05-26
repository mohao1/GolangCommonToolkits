package Stack

/**
根据Stack数据结构的接口规范定义
顺序栈：使用了数组的方式进行数据的存储
T any:定义了泛型的传值可以存储泛型类型的值
通过NewStack构造函数可以获得对应数据实体
*/

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
