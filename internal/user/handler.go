package user

import (
	"RestApi/internal/apperror"
	"RestApi/internal/handlers"
	"RestApi/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersURL = "/users"
	userURL  = "/user/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func New(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetUserList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))

}

func (h *handler) GetUserList(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(200)
	w.Write([]byte(""))
	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(201)
	w.Write([]byte(""))
	return nil
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(200)
	w.Write([]byte(""))
	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte(""))
	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte(""))
	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte(""))
	return nil
}
