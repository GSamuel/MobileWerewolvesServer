package viewmodels

type RoomCreatedResponse struct {
	Code  string `json:"code"`
	Id    string `json:"id"`
	Token string `json:"token"`
}

type JoinRoomResponse struct {
	Code  string `json:"code"`
	Id    string `json:"id"`
	Token string `json:"token"`
}

type RetreiveDataResponse struct {
	Data string `json:"data"`
}

type InfoResponse struct {
}
