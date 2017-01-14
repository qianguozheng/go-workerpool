package workerPool

import "fmt"

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	Len int
}

func NewDispatcher(maxWorkers int) *Dispatcher{
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool:pool, Len:maxWorkers}
}

func (d *Dispatcher) Run()  {
	// starting n number of workers
	fmt.Println("len of workerPool", len(d.WorkerPool))
	for i:=0; i < d.Len; i++{
		fmt.Println("Processor generate worker to do job ", i)
		worker := NewWorker(d.WorkerPool)
		fmt.Println("Generate NewWorker done ", i)
		worker.Start()
		fmt.Println("Worker started", i)
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch()  {
	for{
		select{
		case job:= <-JobQueue:
			// a job request has been received
			fmt.Println("Store a job into jobChannel")
			go func(job Job){
				//try to obtain a worker job channel that is available.
				//this will block until a worker is idle
				jobChannel := <- d.WorkerPool
				//dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

