package player

type Player struct {
	Name string
	Mark string
}

func CreateNewPlayer(Name, Mark string) *Player {
	var newPlayer Player
	newPlayer.Name = Name
	newPlayer.Mark = Mark
	return &newPlayer
}
