package redis

import "github.com/go-redis/redis"

type ClientService interface {

}

type client struct {
	db *redis.Client
}



func New(addr string) ClientService {
	c := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &client{
		db: c,
	}
}
