package viewmodels

type RoomCreatedResponse struct {
	Code string `json:"code"`
}

type JoinRoomResponse struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
}
