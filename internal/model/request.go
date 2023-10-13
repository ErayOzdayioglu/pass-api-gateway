package model

type ServiceEntityRequest struct {
	ServiceName string            `json:"serviceName"`
	IpAddresses []IpAddressEntity `json:"ipAddresses"`
}

type IpAddressEntity struct {
	IpAddress   string `json:"addr"`
	IsAvailable bool   `json:"isAvailable,omitempty"`
}
