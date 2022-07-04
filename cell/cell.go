package cell

type Cell struct {
	Mark string
}

func NewCell() *Cell {
	var c Cell
	c.Mark = ""
	return &c
}

func (c *Cell) IsMarked() bool {
	return c.Mark != ""
}
