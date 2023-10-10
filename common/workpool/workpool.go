package workpool

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	Handler func(v ...interface{})
	Params  []interface{}
}

type Pool struct {
	capacity       uint64
	runningWorkers uint64
	status         int64
	chTask         chan *Task
	sync.Mutex
	PanicHandler func(interface{})
}

var ErrInvalidPoolCap = errors.New("invalid pool cap")

const (
	RUNNING = 1
	STOPED  = 0
)

func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}

	return &Pool{
		capacity: capacity,
		status:   RUNNING,
		// 初始化任务队列
		chTask: make(chan *Task, capacity),
	}, nil
}

func (p *Pool) run() {
	p.runningWorkers++

	go func() {
		defer func() {
			p.runningWorkers--
		}()

		for {
			select {
			case task, ok := <-p.chTask:
				if !ok {
					return
				}
				task.Handler(task.Params...)
			}
		}
	}()

}

func (p *Pool) incRunning() {
	atomic.AddUint64(&p.runningWorkers, 1)
}

func (p *Pool) decRunning() {
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}

func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}

func (p *Pool) GetCap() uint64 {
	return p.capacity
}

func (p *Pool) setStatus(status int64) bool {
	p.Lock()
	defer p.Unlock()

	if p.status == status {
		return false
	}

	p.status = status
	return true
}

var ErrPoolAlreadyClosed = errors.New("pool already closed")

func (p *Pool) Close() {
	p.setStatus(STOPED)

	for len(p.chTask) > 0 {
		time.Sleep(1e6)
	}

	close(p.chTask)
}

func (p *Pool) Put(task *Task) error {
	p.Lock()
	defer p.Unlock()

	if p.GetRunningWorkers() < p.GetCap() {
		p.run()
	}

	if p.status == RUNNING {
		p.chTask <- task
	}

	return nil
}
