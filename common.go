package service

type message struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type registerparams struct {
	Service Service `json:"service"`
}

type findparams struct {
	Name string `json:"name"`
}

// Service represents a service that we want to register. It should hold all information need to be able to be used by another service
type Service struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

type findreply struct {
	ID      string  `json:"id"`
	Jsonrpc string  `json:"jsonrpc"`
	Result  Service `json:"result"`
}
