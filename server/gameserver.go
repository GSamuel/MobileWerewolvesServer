package server

type GameServer struct {
	rooms []*Room
}

func (g *GameServer) CreateRoom() string {
	code := RandomPassword()
	room := NewRoom(code)
	g.rooms = append(g.rooms, room)
	return code
}

func (g *GameServer) JoinRoom(code, nickname string) bool {
	room := g.findRoom(code)

	if room != nil {
		room.JoinPlayer(nickname)
		return true
	}

	return false
	//should probably return errors instead of bools
}

func (g *GameServer) findRoom(code string) *Room {
	for i := 0; i < len(g.rooms); i++ {
		if g.rooms[i].code == code {
			return g.rooms[i]
		}
	}
	return nil
}

func NewGameServer() *GameServer {
	return &GameServer{make([]*Room, 0)}
}
