package job

import "fmt"

var Version = "3.0.5"

type Job struct {
	Name string
}

func (j *Job) Dispatch() {
	fmt.Println("Dispatching job")
}
