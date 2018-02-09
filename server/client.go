package server

type Client struct {
	id       string
	token    string
	nickname string
	master   bool
	data     []Message
}

func (c *Client) AddData(data Message) {
	c.data = append(c.data, data)
}

func (c *Client) RetreiveData() []Message {
	msgs := make([]Message, len(c.data))
	copy(msgs, c.data)
	c.data = nil
	return msgs
}

func NewClient(id, token, nickname string, master bool) *Client {
	return &Client{id, token, nickname, master, make([]Message, 0)}
}
