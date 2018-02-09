package server

import (
	"fmt"
	"github.com/GSamuel/MobileWerewolvesServer/viewmodels"
)

type GameServer struct {
	rooms []*Room
}

func (g *GameServer) CreateRoom() (code, id, token string) {
	c := RandomCode()
	room := NewRoom(c)
	g.rooms = append(g.rooms, room)
	i, t, _ := room.JoinClient("Masterclient", true)
	return c, i, t
}

func (g *GameServer) JoinRoom(code, nickname string) (id, token string, er error) {
	room, err := g.findRoom(code)

	if err != nil {
		return "", "", err
	}

	return room.JoinClient(nickname, false)
	//should return error when room does not exist
	//and when nickname already exists in room
}

func (g *GameServer) RoomInfo(code string) (viewmodels.Room, error) {
	r, err := g.findRoom(code)

	if err != nil {
		return viewmodels.Room{}, err
	}

	return r.Info(), nil
}

func (g *GameServer) findRoom(code string) (*Room, error) {
	for i := 0; i < len(g.rooms); i++ {
		if g.rooms[i].code == code {
			return g.rooms[i], nil
		}
	}
	return nil, fmt.Errorf("Room with code %s does not exist", code)
}

func (g *GameServer) SendData(code string, target int, method int, data string) error {
	r, err := g.findRoom(code)

	if err != nil {
		return err
	}

	r.SendData(code, target, method, data)
	return nil
}

func (g *GameServer) RetreiveData(code string) (string, error) {
	r, err := g.findRoom(code)

	if err != nil {
		return "", err
	}

	return r.RetreiveData(code), nil
}

func NewGameServer() *GameServer {
	return &GameServer{make([]*Room, 0)}
}
