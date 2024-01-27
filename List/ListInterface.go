package List

type Interface[T any] interface {
	Get(index int) T
	Add(val T) int
	Remove(index int) T
	UpData(index int, val T) bool
	Insert(index int, val T) bool
	GetLength() int
	IsEmpty() bool
	IsNoEmpty() bool
}
