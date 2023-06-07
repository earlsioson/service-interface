package calculator

var Version = "3.0.0"

type Calculator struct {
	Name string
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}
