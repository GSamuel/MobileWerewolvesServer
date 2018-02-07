package main

import (
	"fmt"
	"github.com/GSamuel/MobileWerewolvesServer/server"
	"github.com/GSamuel/MobileWerewolvesServer/utils"
	"github.com/GSamuel/MobileWerewolvesServer/viewmodels"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var gameServer server.GameServer

func httpHandler(w http.ResponseWriter, r *http.Request) {
}

func createRoom(w http.ResponseWriter, req *http.Request) {
	//check if room already exists for this connection.
	code := gameServer.CreateRoom()

	response := viewmodels.RoomCreatedResponse{code}

	utils.Write(w, response)

	//create a new room with unique id.
	//send id back to room requester.
	//send unique connection secret to requester. (authorization token)
}

func joinRoom(w http.ResponseWriter, r *http.Request) {
	var t viewmodels.JoinRoomRequest
	utils.Read(r, &t)

	fmt.Println(t.Code, t.Nickname)

	success := gameServer.JoinRoom(t.Code, t.Nickname)

	response := viewmodels.JoinRoomResponse{t.Code, success}
	utils.Write(w, response)
}

func main() {
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/", httpHandler) // set router
	http.HandleFunc("/create", createRoom)
	http.HandleFunc("/join", joinRoom)

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
