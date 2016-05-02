package service

type whereis_request struct {
	Jsonrpc     string `json:"jsonrpc"`
	ID          string `json:"id"`
	Method      string `json:"method"`
	ServiceName string `json:"params"`
}

type register_request struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type registerparams struct {
	Service Service `json:"service"`
}

// Service represents a service that we want to register. It should hold all information need to be able to be used by another service
type Service struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

type whereis_reply struct {
	ID      string  `json:"id"`
	Jsonrpc string  `json:"jsonrpc"`
	Result  Service `json:"result"`
}
