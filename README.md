# CubeGo

## 介绍
Golang常用数据扩展工具包。

## 项目地址
- GitHub:https://github.com/mohao1/CubeGo
- Gitee:https://gitee.com/mohaos/CubeGo

## 功能设计

### 数据结构模块
**模块设计**
1. Heap 堆
   - 大顶堆：maxHeap。
   - 小顶堆：minHeap。
2. Stack 栈
   - Stack：顺序栈。
   - LinkStack：链表栈
3. Queue 队列
   - Queue：顺序队列
   - LinkQueue：链表队列
4. List：链表
   - LinkList：单向链表
   - BLinkList：双向链表

### 其他工具实现
**模块设计**
1. Util 工具
   - Compare：选择二数选一，选择大/小。
   - Equal：数据比较，比较数据大小。
2. ExecutorsPool：线程
   - Executors：线程池

### Redis工具实现
**模块设计**
1. RedisQueue：Redis实现队列
2. RedisLock：Redis实现的分布式锁
   - 互斥锁
   - 读写锁 = 等待更新

### 日志工具
**模块设计**
1. Logger日志
   - logs
   - logx

### ORM快速代码生成工具
**模块设计**
1. XGen 基于Gorm的代码生成实现
   - 通过模板实现Gorm的操作简易封装


**持续更新**

