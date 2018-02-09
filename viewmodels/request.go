package viewmodels

type JoinRoomRequest struct {
	Code     string `json:"code"`
	Nickname string `json:"nickname"`
}

type SendDataRequest struct {
	Code   string `json:"code"`
	Id     string `json:"id"`
	Token  string `json:"token"`
	Target string `json:"target"`
	Data   string `json:"data"`
}

type RetreiveDataRequest struct {
	Code  string `json:"code"`
	Id    string `json:"id"`
	Token string `json:"token"`
}

type InfoRequest struct {
	Code string `json:"code"`
}
