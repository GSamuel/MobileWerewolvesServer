package server

import (
	"fmt"
	"github.com/GSamuel/MobileWerewolvesServer/viewmodels"
	"sync"
)

type Room struct {
	code    string
	clients []*Client
	data    []string
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

func (r *Room) SendData(code string, target int, method int, data string) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if method == 0 {
		for i := 0; i < len(r.clients); i++ {
			r.clients[i].AddData(data)
		}
	}
	if method == 1 {
		//send to target
		r.clients[target].AddData(data)
	}
	if method == 2 {
		r.data = append(r.data, data)
	}
}

func (r *Room) RetreiveData(code string) string {
	r.mux.Lock()
	defer r.mux.Unlock()

	if len(r.data) == 0 {
		return ""
	}

	return r.data[len(r.data)-1]
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
	return &Room{code, make([]*Client, 0), make([]string, 0), sync.Mutex{}}
}
