package task2

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type Result struct {
	TaskID   int
	Duration time.Duration
	Error    error
}

type Scheduler struct {
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Run(task []Task) []Result {
	const (
		defaultMaxConcurrency = 20
	)

	var (
		wg  sync.WaitGroup
		res = make([]Result, len(task))
		mu  sync.Mutex
		sem = make(chan struct{}, defaultMaxConcurrency)
	)

	for taskID, task := range task {
		wg.Add(1)
		sem <- struct{}{}
		go func(id int, t Task) {
			defer func() {
				<-sem
				wg.Done()
			}()

			start := time.Now()

			var ret Result

			defer func() {
				if e := recover(); e != nil {
					ret = Result{
						TaskID:   id,
						Duration: time.Since(start),
						Error:    fmt.Errorf("%v", e),
					}
				}
			}()

			t()

			ret = Result{
				TaskID:   id,
				Duration: time.Since(start),
				Error:    nil,
			}
			mu.Lock()
			res[id] = ret
			mu.Unlock()

		}(taskID, task)
	}

	wg.Wait()
	return res
}
