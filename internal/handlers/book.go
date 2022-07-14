package handlers

import (
	"RestApi/internal/apperror"
	"RestApi/internal/book"
	"RestApi/pkg/logging"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

const (
	bookURL        = "/books/:id"
	booksURL       = "/books"
	booksCreateURL = "/bookCreate"
	booksUpdateURL = "/bookUpdate"
)

type handler struct {
	logger     *logging.Logger
	repository book.Repository
}

func NewHandler(repository book.Repository, logger *logging.Logger) Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, booksURL, apperror.Middleware(h.FindAll))
	router.HandlerFunc(http.MethodGet, bookURL, apperror.Middleware(h.FindOne))
	router.HandlerFunc(http.MethodPost, booksCreateURL, apperror.Middleware(h.Create))
	router.HandlerFunc(http.MethodPost, booksUpdateURL, apperror.Middleware(h.Update))
	router.HandlerFunc(http.MethodDelete, bookURL, apperror.Middleware(h.Delete))
}

func (h *handler) FindAll(w http.ResponseWriter, r *http.Request) error {
	all, err := h.repository.FindAll(context.TODO())
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}

func (h *handler) FindOne(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	oneBook, err := h.repository.FindOne(context.TODO(), id)
	if err != nil {
		return err
	}
	allBytes, err := json.Marshal(oneBook)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)
	return nil
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) error {
	var d book.Book
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return err
	}
	bookId, err := h.repository.Create(context.TODO(), d)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(bookId))
	return nil
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) error {
	var d book.Book
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return err
	}
	err := h.repository.Update(context.TODO(), d)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	err := h.repository.Delete(context.TODO(), id)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}
