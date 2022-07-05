package cell

import "strconv"

type Cell struct {
	Mark string
}

var markedLocations []int

func NewCell(postionValue int) *Cell {
	var c Cell
	c.Mark = strconv.Itoa(postionValue)
	return &c
}

func IsMarked(PlayerMaark string, position int) bool {
	for i := 0; i < len(markedLocations); i++ {
		if position == markedLocations[i] {
			return true

		}

	}
	return false
}
