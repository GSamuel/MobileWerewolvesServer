package server

type Room struct {
	code    string
	players []*Player
}

func (r *Room) JoinPlayer(nickname string) {
	//should check that player does not exist already
	r.players = append(r.players, NewPlayer(nickname))
}

func (r *Room) PlayerCount() int {
	return len(r.players)
}

func NewRoom(code string) *Room {
	return &Room{code, make([]*Player, 0)}
}
