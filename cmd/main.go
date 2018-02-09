package main

import (
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

func sendData(w http.ResponseWriter, r *http.Request) {
	var t viewmodels.SendDataRequest
	utils.Read(r, &t)

	gameServer.SendData(t.Code, t.Target, t.Method, t.Data)
}

func retreiveData(w http.ResponseWriter, r *http.Request) {
	var t viewmodels.RetreiveDataRequest
	utils.Read(r, &t)
	data, err := gameServer.RetreiveData(t.Code)

	if err != nil {
		//do something here
	}

	response := viewmodels.RetreiveDataResponse{data}
	utils.Write(w, response)
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	//check if room already exists for this connection.
	code, id, token := gameServer.CreateRoom()

	response := viewmodels.RoomCreatedResponse{code, id, token}

	utils.Write(w, response)

	//create a new room with unique id.
	//send id back to room requester.
	//send unique connection secret to requester. (authorization token)
}

func joinRoom(w http.ResponseWriter, r *http.Request) {
	var t viewmodels.JoinRoomRequest
	utils.Read(r, &t)
	id, token, err := gameServer.JoinRoom(t.Code, t.Nickname)

	if err != nil {
		return
	}

	response := viewmodels.JoinRoomResponse{t.Code, id, token}
	utils.Write(w, response)
}

func info(w http.ResponseWriter, r *http.Request) {
	var t viewmodels.InfoRequest
	utils.Read(r, &t)

	info, err := gameServer.RoomInfo(t.Code)

	if err != nil {
		return
	}

	utils.Write(w, info)
}

func main() {
	rand.Seed(time.Now().Unix())

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/create", createRoom)
	http.HandleFunc("/join", joinRoom)
	http.HandleFunc("/send", sendData)
	http.HandleFunc("/retreive", retreiveData)
	http.HandleFunc("/info", info)

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
