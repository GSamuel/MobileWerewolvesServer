package server

import (
	"fmt"
	"github.com/GSamuel/MobileWerewolvesServer/viewmodels"
	"sync"
)

type Room struct {
	code    string
	clients []*Client
	mux     sync.Mutex
}

func (r *Room) JoinClient(nickname string, master bool) (string, string, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	for i := 0; i < len(r.clients); i++ {
		if r.clients[i].nickname == nickname {
			return "", "", fmt.Errorf("Nickname %s is already in use in room with code %s", nickname, r.code)
		}

	}

	client := NewClient(RandomId(), RandomId(), nickname, master)
	r.clients = append(r.clients, client)
	return client.id, client.token, nil
}

func (r *Room) ClientCount() int {
	r.mux.Lock()
	defer r.mux.Unlock()

	return len(r.clients)
}

func (r *Room) SendData(id, token, target, data string) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	sender, err := r.findClient(id)

	if err != nil {
		return err
	}

	if sender.token != token {
		return fmt.Errorf("Invalid token for client with id %s", id)
	}

	receiver, err := r.findClient(target)

	if err != nil {
		return err
	}

	receiver.AddData(Message{id, data})
	return nil
}

func (r *Room) RetreiveData(id, token string) ([]Message, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	requester, err := r.findClient(id)

	if err != nil {
		return []Message{}, err
	}

	if requester.token != token {
		return []Message{}, fmt.Errorf("Invalid token for client with id %s", id)
	}

	return requester.RetreiveData(), nil
}

func (r *Room) findClient(id string) (*Client, error) {
	for i := 0; i < len(r.clients); i++ {
		if r.clients[i].id == id {
			return r.clients[i], nil
		}
	}
	return nil, fmt.Errorf("Client with id %s does not exist", id)
}

func (r *Room) Info() viewmodels.Room {
	r.mux.Lock()
	defer r.mux.Unlock()

	clients := make([]viewmodels.Client, 0)

	for i := 0; i < len(r.clients); i++ {
		c := r.clients[i]
		clients = append(clients, viewmodels.Client{c.id, c.nickname, c.master})
	}

	return viewmodels.Room{r.code, clients}
}

func NewRoom(code string) *Room {
	return &Room{code, make([]*Client, 0), sync.Mutex{}}
}
