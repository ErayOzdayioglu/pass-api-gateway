package model

type AddToServiceRegistryRequest struct {
	ServiceName string `json:"serviceName"`
	IpAddress   string `json:"ipAddress"`
	Port        int    `json:"port"`
}
