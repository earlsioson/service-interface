package calculator

var Version = "2.0.2"

type Calculator struct {
	Name string
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}
