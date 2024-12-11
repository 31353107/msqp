package tasks

import (
	"sync"
	"time"
)

// Task 定时任务结构体
type Task struct {
	ticker   *time.Ticker
	stopChan chan struct{}
	wg       sync.WaitGroup
}

// NewTask 创建一个新的定时任务
func NewTask(interval time.Duration, job func()) *Task {
	task := &Task{
		ticker:   time.NewTicker(interval),
		stopChan: make(chan struct{}),
	}

	task.wg.Add(1)
	go func() {
		defer task.wg.Done()
		for {
			select {
			case <-task.ticker.C:
				job() // 执行定时任务
			case <-task.stopChan:
				return // 停止任务
			}
		}
	}()
	return task
}

// Stop 停止定时任务
func (t *Task) Stop() {
	t.ticker.Stop()
	close(t.stopChan)
	t.wg.Wait()
}
