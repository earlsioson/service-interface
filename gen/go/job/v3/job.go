package job

import "fmt"

var Version = "3.0.6"

type Job struct {
	Name string
}

func (j *Job) Dispatch() {
	fmt.Println("Dispatching job")
}
