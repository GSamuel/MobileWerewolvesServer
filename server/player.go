package server

type Player struct {
	nickname string
}

func NewPlayer(nickname string) *Player {
	return &Player{nickname}
}
