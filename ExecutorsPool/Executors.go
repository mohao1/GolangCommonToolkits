package ExecutorsPool

import "sync"

/**
1、创建协程池
2、需要实现Executors的Interface接口的GoRun方法
3、协程池中含有Wg可以控制main协程的同步阻塞等待
*/

// Executors 线程池
type Executors struct {
	maxSize   int              //最大大小
	typeSize  int              //轮询记录任务分配数量
	taskQueue []chan Interface //任务队列集合
	wg        sync.WaitGroup   //阻塞wg
}

// Run 执行任务
func (e *Executors) Run(task Interface) {
	switch e.maxSize {
	case 1:
		{
			e.wg.Add(1)
			e.taskQueue[0] <- task
		}
	case e.typeSize:
		{
			e.wg.Add(1)
			e.typeSize = 0
			e.taskQueue[e.typeSize] <- task

		}

	default:
		e.wg.Add(1)
		e.taskQueue[e.typeSize] <- task
		e.typeSize++
	}
}

// StartWorkerPool 启动线程池
func (e *Executors) StartWorkerPool() {
	for i := 0; i < e.maxSize; i++ {
		e.coreTaskQueue(make(chan Interface, 10000))
		go e.startWorker(i)
	}
}

// 设置chan
func (e *Executors) coreTaskQueue(taskQueue chan Interface) {
	e.taskQueue = append(e.taskQueue, taskQueue)
}

// 设置线程池的大小
func (e *Executors) coreMaxPoolSize(size int) {
	e.maxSize = size
}

func (e *Executors) GetWg() *sync.WaitGroup {
	return &e.wg
}

// 启动线程
func (e *Executors) startWorker(index int) {
	for {
		select {
		case task := <-e.taskQueue[index]:
			{
				task.GoRun()
				e.wg.Done()
			}
		}
	}
}

func NewExecutors(size int) *Executors {
	return &Executors{
		maxSize:   size,
		typeSize:  0,
		taskQueue: make([]chan Interface, 0),
		wg:        sync.WaitGroup{},
	}
}
