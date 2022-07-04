package player

import (
	"fmt"
	"testing"
)

var testdataforPlayer = []struct {
	FirstName    string
	Mark         string
	ExpectedName string
	ExpectedMark string
}{
	{"Darshan", "X", "Darshan", "X"}, {"Rahul", "O", "Rahul", "O"}, {"Prathamesh", "X", "Prathamesh", "X"}, {"Himanshu", "O", "Himanshu", "O"}, {"Prathamesh", "X", "Prathamesh", "X"},
}

func TestCreatePlayer(t *testing.T) {

	for i := 0; i < len(testdataforPlayer); i++ {
		var newPlayer Player
		newPlayer.Name = testdataforPlayer[i].FirstName
		newPlayer.Mark = testdataforPlayer[i].Mark
		if newPlayer.Name == testdataforPlayer[i].ExpectedName && newPlayer.Mark == testdataforPlayer[i].ExpectedMark {
			fmt.Println("Test Pass Sucessfully")
		} else {
			fmt.Println("Test Not Pass")
		}

	}
}
