package calculator

<<<<<<< Updated upstream
var Version = "2.0.3"
=======
var Version = "2.0.4"
>>>>>>> Stashed changes

type Calculator struct {
	Name string
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}
