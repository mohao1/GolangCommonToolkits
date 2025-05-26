# Golang Common Toolkits

## Introduction
Common data extension toolkits for Golang.

## Project Links
- GitHub: https://github.com/mohao1/GolangCommonToolkits?tab=readme-ov-file
- Gitee: https://gitee.com/mohaos/golang-common-toolkits

## Functional Design

### Data Structure Module
**Module Design**
1. Heap
   - Max Heap: maxHeap.
   - Min Heap: minHeap.
2. Stack
   - Stack: Sequential stack.
   - LinkStack: Linked list stack.
3. Queue
   - Queue: Sequential queue.
   - LinkQueue: Linked list queue.
4. List: Linked list
   - LinkList: Singly linked list.
   - BLinkList: Doubly linked list.
   
### Other Tool Implementations
**Module Design**
1. Util Tools
   - Compare: Choose between two numbers, select the larger/smaller one.
   - Equal: Data comparison, compare data sizes.
2. ExecutorsPool: Thread
   - Executors: Thread pool.

### Redis Tool Implementations
**Module Design**
1. RedisQueue: Queue implemented with Redis.
2. RedisLock: Distributed locks implemented with Redis.
   - Mutex lock.
   - Read-write lock = Waiting for updates.

### Logging Tools
**Module Design**
1. Logger Logging
- logs
- logx

### ORM Rapid Code Generation Tool
**Module Design**
1. XGen: Code generation implementation based on Gorm.
- Simplify Gorm operations through template-based encapsulation.

**Continuously Updated**