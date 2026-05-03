package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type RedisClient interface {
	Set(key string, value string) bool
}

type RedisService struct {
	client RedisClient
}

func NewRedisService(client RedisClient) *RedisService {
	return &RedisService{
		client: client,
	}
}

func (r *RedisService) Register(url URLRequest) (urlKey string, err error) {
	sha := sha256.Sum256([]byte(url.Value))
	shaString := hex.EncodeToString(sha[:])
	ok := r.client.Set(shaString, url.Value)
	if !ok {
		return "", errors.New("failed to store url in datastore")
	}
	return shaString, nil
}
