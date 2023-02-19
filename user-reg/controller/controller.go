package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"user-reg/helper"
	"user-reg/model"
	"user-reg/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Controller interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

func New(svc service.Service, validate *validator.Validate) Controller {
	return &ctl{
		svc:      svc,
		validate: validate,
	}
}

type ctl struct {
	svc      service.Service
	validate *validator.Validate
}

func (c *ctl) CreateUser(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	if header != "application/json" {
		c.writeJSON(w, http.StatusUnsupportedMediaType, model.Response{
			Status:  "error",
			Message: "Content-Type header is not application/json",
		})
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	var user model.CreateUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		c.writeJSON(w, http.StatusBadRequest, model.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if err = c.validate.Struct(user); err != nil {
		e := err.(validator.ValidationErrors)
		c.writeJSON(w, http.StatusBadRequest, model.Response{
			Status:  "error",
			Message: fmt.Sprintf("bad %s", e[0].Tag()),
		})
		return
	}

	if err = c.svc.CreateUser(r.Context(), &user); err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, model.ErrTimeOut) {
			c.writeJSON(w, http.StatusGatewayTimeout, model.Response{
				Status:  "timeout",
				Message: err.Error(),
			})
			return
		}
		if errors.Is(err, model.ErrUserExist) {
			c.writeJSON(w, http.StatusBadRequest, model.Response{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}
		c.writeJSON(w, http.StatusInternalServerError, model.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.writeJSON(w, http.StatusCreated, model.Response{
		Status:  "ok",
		Message: "user created",
	})
}

func (c *ctl) GetUser(w http.ResponseWriter, r *http.Request) {
	e := chi.URLParam(r, "email")
	if !helper.ValidateEmail(e) {
		c.writeJSON(w, http.StatusBadRequest, model.Response{
			Status:  "error",
			Message: "email not valid",
		})
		return
	}
	user, err := c.svc.GetUserByEmail(r.Context(), e)
	if err != nil && errors.Is(err, context.DeadlineExceeded) {
		c.writeJSON(w, http.StatusGatewayTimeout, model.Response{
			Status:  "timeout",
			Message: err.Error(),
		})
		return
	}
	if err != nil {
		c.writeJSON(w, http.StatusUnauthorized, model.Response{
			Status:  "not found",
			Message: err.Error(),
		})
		return
	}
	c.writeJSON(w, http.StatusOK, user)
}

func (c *ctl) writeJSON(w http.ResponseWriter, status int, respBody any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(respBody)
}
