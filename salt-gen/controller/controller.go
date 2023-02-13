package controller

import (
	"encoding/json"
	"net/http"
	"salt-gen/service"
)

type Controller interface {
	SaltGen(w http.ResponseWriter, r *http.Request)
}

func New(length int) Controller {
	return &ctl{length: length}
}

type ctl struct{
	length int
}

func (c *ctl) SaltGen(w http.ResponseWriter, r *http.Request) {

	salt, err := service.Generate(c.length)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	data := struct {
		Salt string `json:"salt"`
	}{
		Salt: salt,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
