package Queue

/**
根据Queue数据结构的接口规范定义
顺序队列：使用了数组的方式进行数据的存储
T any:定义了泛型的传值可以存储泛型类型的值
通过NewQueue构造函数可以获得对应数据实体
*/

// Queue 顺序队列
type Queue[T any] struct {
	data []T
	len  int
	size int
}

func (q *Queue[T]) Push(value T) bool {
	if q.len == q.size {
		return false
	} else {
		q.data = append(q.data, value)
		q.len++
		return true
	}
}

func (q *Queue[T]) Pop() T {
	if q.IsEmpty() {
		return *new(T)
	} else {
		data := q.data[0]
		q.data = q.data[1:]
		q.len--
		return data
	}
}

func (q *Queue[T]) Size() int {
	return q.len
}

func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		return *new(T)
	} else {
		return q.data[0]
	}
}

func (q *Queue[T]) Back() T {
	if q.IsEmpty() {
		return *new(T)
	} else {
		return q.data[q.len-1]
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue[T]) IsNoEmpty() bool {
	return !q.IsEmpty()
}

// NewQueue 构造函数
func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0, size),
		len:  0,
		size: size,
	}
}
