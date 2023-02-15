package model

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
