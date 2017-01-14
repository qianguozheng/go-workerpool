package workerPool

import (
	"fmt"
)

type Producer struct {
	job Job
}

func NewProducer(maxJob int) (*Producer) {
	job := Job{Payload:Payload(maxJob)}
	return &Producer{job: job}
}

func (p Producer) Run() {
	for i:=1; i < int(p.job.Payload); i++{
		work:= Job{Payload: Payload(i)}
		fmt.Println("Producer: job " , i)
		JobQueue <- work
		//time.Sleep(time.Second*1)
	}
}


//func payloadHandler(){
//	//Go through each payload and queue items individually to be posted to S3
//	//for _, payload := range
//	payload := Payload(4)
//	work := Job{Payload: payload}
//
//	//Push the work onto the queue.
//	JobQueue <- work
//}