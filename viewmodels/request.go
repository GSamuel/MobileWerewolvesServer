package viewmodels

type JoinRoomRequest struct {
	Code     string `json:"code"`
	Nickname string `json:"nickname"`
}

type SendDataRequest struct {
	Code   string `json:"code"`
	Target int    `json:"target"`
	Method int    `json:"method"` //All clients, Server, Specific client
	Data   string `json:"data"`
}

type RetreiveDataRequest struct {
	Code string `json:"code"`
}

type InfoRequest struct {
	Code string `json:"code"`
}
