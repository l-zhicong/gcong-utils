package cpool

import (
	"container/list"
	"sync"
)

type TaskList struct {
	rwLock *sync.RWMutex //读写锁
	Job    *list.List    //链表
}

func NewList() *TaskList {
	return &TaskList{
		Job:    list.New(),
		rwLock: new(sync.RWMutex),
	}
}

func (t TaskList) PushTask(f any) {
	defer t.rwLock.Unlock()
	t.rwLock.Lock()
	t.Job.PushFront(f) //头部插入
}

//取链表最后一个删除

func (t TaskList) PopBack() (value interface{}) {
	defer t.rwLock.Unlock()
	t.rwLock.Lock()
	v := t.Job.Back() //取最后一个
	if v != nil {
		value = t.Job.Remove(v) //删除
		return
	}
	return
}
