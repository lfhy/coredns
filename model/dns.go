package model

type Dns struct {
	Origin     string
	NameServer string
	TTL        int
	Key        string
	Value      string
}

type A struct {
	Host string `json:"host,omitempty"`
	TTL  int    `json:"ttl,omitempty"`
}

type DnsData struct {
	Data []*Dns `json:"data"`
}

type Message struct {
	Error string
}

// DNS操作消息
var NewMessage = make(chan []*Dns, 1)
