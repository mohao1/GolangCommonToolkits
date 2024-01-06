package Queue

/**
Queue数据结构的接口规范定义
T any:定义了泛型的传值可以存储泛型类型的值
定义了:
	Push(value T) bool //添加数据
	Pop() T            //弹出数据
	Size() int         //查看个数
	Peek() T           //查看队列头部的第一个元素
	Back() T           //获取队列尾部的第一个元素
	IsEmpty() bool     //是否队列空
	IsNoEmpty() bool   //是否队列不为空
*/

type Interface[T any] interface {
	Push(value T) bool //添加数据
	Pop() T            //弹出数据
	Size() int         //查看个数
	Peek() T           //查看队列头部的第一个元素
	Back() T           //获取队列尾部的第一个元素
	IsEmpty() bool     //是否队列空
	IsNoEmpty() bool   //是否队列不为空
}
