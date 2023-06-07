package calculator

var Version = "1.0.0"

type Calculator struct {
	Name string
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}
