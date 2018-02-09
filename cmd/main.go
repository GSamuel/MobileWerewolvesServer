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

func sendData(w http.ResponseWriter, r *http.Request) {
	var t viewmodels.SendDataRequest
	utils.Read(r, &t)

	err := gameServer.SendData(t.Code, t.Id, t.Token, t.Target, t.Data)
	if err != nil {
		fmt.Println(err)
	}
}

func retreiveData(w http.ResponseWriter, r *http.Request) {
	var t viewmodels.RetreiveDataRequest
	utils.Read(r, &t)
	data, err := gameServer.RetreiveData(t.Code, t.Id, t.Token)

	if err != nil {
		fmt.Println(err)
		return
	}

	msgs := make([]viewmodels.Message, 0)

	for i := 0; i < len(data); i++ {
		msgs = append(msgs, viewmodels.Message{data[i].Id, data[i].Data})
	}

	response := viewmodels.RetreiveDataResponse{msgs}
	utils.Write(w, response)
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	code, id, token := gameServer.CreateRoom()

	response := viewmodels.RoomCreatedResponse{code, id, token}

	utils.Write(w, response)
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
