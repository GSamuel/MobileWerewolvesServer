package viewmodels

type Room struct {
	Code    string   `json:"code"`
	Clients []Client `json:"clients"`
}

type Client struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Master   bool   `json:"master"`
}

type Message struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}
