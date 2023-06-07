package job

import "fmt"

var Version = "1.0.0"

type Job struct {
	Name string
}

func (j *Job) Dispatch() {
	fmt.Println("Dispatching job")
}
