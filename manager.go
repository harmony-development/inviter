package main

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Manager struct {
	Connections *lru.Cache
}

func NewManager() (*Manager, error) {
	manager := &Manager{}

	cache, err := lru.NewWithEvict(4096, manager.OnConnectionEvict)
	if err != nil {
		return nil, err
	}
	manager.Connections = cache
	return manager, err
}

func (im Manager) Connect(host string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := grpc.DialContext(ctx, host, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	if err != nil {
		return nil, err
	}
	im.Connections.Add(host, client)
	return client, nil
}

func (im Manager) GetOrConnect(host string) (*grpc.ClientConn, error) {
	conn, exists := im.Connections.Get(host)
	if exists {
		return conn.(*grpc.ClientConn), nil
	} else {
		return im.Connect(host)
	}
}

func (im Manager) OnConnectionEvict(key, value interface{}) {
	conn := value.(*grpc.ClientConn)
	if err := conn.Close(); err != nil {
		log.Println(err)
	}
}
