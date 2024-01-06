package LinkQueue

type Node[T any] struct {
	Data    T
	curNode *Node[T]
}

/**
根据Queue数据结构的接口规范定义
顺序队列：使用了数组的方式进行数据的存储
T any:定义了泛型的传值可以存储泛型类型的值
通过NewLinkQueue构造函数可以获得对应数据实体
*/

// LinkQueue 链表队列
type LinkQueue[T any] struct {
	topData *Node[T]
	endData *Node[T]
	len     int
	size    int
}

func (l *LinkQueue[T]) Push(value T) bool {
	if l.size == l.len {
		return false
	} else {
		l.endData.curNode = &Node[T]{
			Data:    value,
			curNode: nil,
		}
		l.endData = l.endData.curNode
		l.len++
		return true
	}
}

func (l *LinkQueue[T]) Pop() T {
	if l.IsEmpty() {
		return *new(T)
	} else {
		data := l.topData.curNode.Data
		l.topData.curNode = l.topData.curNode.curNode
		l.len--
		if l.len == 0 {
			l.endData = l.topData
		}
		return data
	}
}

func (l *LinkQueue[T]) Size() int {
	return l.len
}

func (l *LinkQueue[T]) Peek() T {
	if l.IsEmpty() {
		return *new(T)
	} else {
		return l.topData.curNode.Data
	}
}

func (l *LinkQueue[T]) Back() T {
	if l.IsEmpty() {
		return *new(T)
	} else {
		return l.endData.curNode.Data
	}
}

func (l *LinkQueue[T]) IsEmpty() bool {
	return l.len == 0
}

func (l *LinkQueue[T]) IsNoEmpty() bool {
	return !l.IsEmpty()
}

func NewLinkQueue[T any](size int) *LinkQueue[T] {
	data := &Node[T]{}
	return &LinkQueue[T]{
		topData: data,
		endData: data,
		size:    size,
		len:     0,
	}
}
