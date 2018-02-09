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
	Messages []Message `json:"messages"`
}
