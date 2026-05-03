package main

type RedisClient interface {
	Set(key string, value string) string
}

type RedisService struct {
	client RedisClient
}

func NewRedisService(client RedisClient) *RedisService {
	return &RedisService{
		client: client,
	}
}

func (r *RedisService) Register(url URL) (urlKey string, err error) {
	uuid := r.client.Set(url.Key, url.Value)
	return uuid, nil
}
