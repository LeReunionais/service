package service

type message struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method 	string `json:"method"`
	Params  params `json:"params"`
}

type params struct {
	Service Service `json:"service"`
}

// Service represents a service that we want to register. It should hold all information need to be able to be used by another service
type Service struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}
