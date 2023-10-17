package service

type AddServiceRequest struct {
	ServiceName string `json:"serviceName"`
	Path        string `json:"path"`
	Hosts       []Host `json:"hosts"`
}
