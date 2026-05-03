package main

import (
	"net/http"
)

type URL struct {
	Key   string
	Value string
}

type URLService interface {
	Register(url URL) (urlKey string, err error)
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
	urlKey, err := u.service.Register(URL{Key: r.URL.Query().Get("key")})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(urlKey))
}
