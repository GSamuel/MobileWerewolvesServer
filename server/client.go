package server

type Client struct {
	id       string
	token    string
	nickname string
	master   bool
	data     []string
}

func (c *Client) AddData(data string) {
	c.data = append(c.data, data)
}

func NewClient(id, token, nickname string, master bool) *Client {
	return &Client{id, token, nickname, master, make([]string, 0)}
}
