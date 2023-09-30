package model

type GetServiceRegistryResponse struct {
	ServiceName string `json:"serviceName"`
	IpAddress   string `json:"ipAddress"`
	Port        int    `json:"port"`
}
