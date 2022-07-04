package cell

import (
	"fmt"
	"testing"
)

var testdataforcellMark = []struct {
	InputString    string
	ExpectedString string
}{
	{"", ""}, {"", ""}, {"", ""}, {"", ""}, {"", ""},
}

func TestIsMarked(t *testing.T) {
	for i := 0; i < len(testdataforcellMark); i++ {
		if testdataforcellMark[i].ExpectedString == testdataforcellMark[i].InputString {
			fmt.Println("Data pass")

		}
	}
}
