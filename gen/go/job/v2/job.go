package job

import "fmt"

var Version = "2.0.3"

type Job struct {
	Name string
}

func (j *Job) Dispatch() {
	fmt.Println("Dispatching job")
}
