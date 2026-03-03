package core

import "sync"

type Job func()

type WorkerPool struct {
	jobs chan Job
	wg   sync.WaitGroup
}

func NewWorkerPool(size int) *WorkerPool {
	wp := &WorkerPool{
		jobs: make(chan Job),
	}
	for i := 0; i < size; i++ {
		go func() {
			for job := range wp.jobs {
				job()
				wp.wg.Done()
			}
		}()
	}
	return wp
}

func (wp *WorkerPool) Submit(job Job) {
	wp.wg.Add(1)
	wp.jobs <- job
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}