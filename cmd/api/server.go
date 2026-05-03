package main

import (
	"net/http"
)

type URL struct {
	Key   string
	Value string
}

type URLRequest struct {
	Value string
}

type URLService interface {
	Register(url URLRequest) (urlKey string, err error)
}

type UrlServer struct {
	service URLService
}

func NewUrlServer(service URLService) *UrlServer {
	return &UrlServer{
		service: service,
	}
}

func (u *UrlServer) RegisterURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	value := r.URL.Query().Get("value")
	if value == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	urlKey, err := u.service.Register(URLRequest{Value: value})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(urlKey))
}
