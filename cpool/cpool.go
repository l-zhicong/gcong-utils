package cpool

//TODO 等待子协程完事sync.WaitGroup

type Pool struct {
	max      int   //最大携程数
	count    Count //正在运行任务数量
	taskList *TaskList
	closed   *bool //是否关闭
}

var pool = New()

func New() *Pool {
	p := &Pool{
		max:      0,
		taskList: NewList(),
		count:    DefaultCount(),
		closed:   DefaultClosed(),
	}
	return p
}

func (p *Pool) SetConfig(max int) *Pool {
	if max > 0 {
		p.max = max
	}
	return p
}

func DefaultClosed() *bool {
	defaultClosed := true
	return &defaultClosed
}

//加入任务列表
//任务数量大于设置数量

func (p *Pool) AddJob(f any) error {
	p.taskList.PushTask(f)

	for {
		c := int(p.count)
		if p.max != 0 && c >= p.max { //防止无限循环,限制最大携程
			return nil
		}
		if p.count.Cas(c, c+1) { //现有携程 等于 现有携程+1 然后替换现有携程数
			break
		}
	}
	p.fork()
	return nil
}

func (p *Pool) fork() {
	go func() {
		defer p.count.Add(-1)
		for *p.closed {
			if Job := p.taskList.PopBack(); Job != nil {
				Job.(func())()
			} else {
				return
			}
		}
	}()
}

func (p *Pool) GetCount() int64 {
	return *p.count.int64()
}

func (p *Pool) Closed() {
	closed := false
	p.closed = &closed
}
