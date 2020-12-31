package manager

import (
	"context"
	"crypto/tls"
	"log"
	"strings"
	"time"

	chatv1 "github.com/harmony-development/inviter/gen/chat/v1"
	lru "github.com/hashicorp/golang-lru"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Manager manages the connections to various Harmony homeservers
type Manager struct {
	Connections *lru.Cache
}

// New creates a new Manager
func New() *Manager {
	manager := &Manager{}

	cache, err := lru.NewWithEvict(4096, manager.onConnectionEvict)
	if err != nil {
		panic(err)
	}

	manager.Connections = cache
	return manager
}

func (im Manager) connect(host string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := grpc.DialContext(ctx, host, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	if err != nil {
		return nil, err
	}
	im.Connections.Add(host, client)
	return client, nil
}

// InviteData holds the invite data of a server
type InviteData struct {
	GuildName   string
	Avatar      string
	AvatarHost  string
	MemberCount int
	Host        string
	InviteID    string
}

// connectOrCached gets a connection to a host
func (im Manager) connectOrCached(host string) (*grpc.ClientConn, error) {
	if !strings.Contains(host, ":") {
		host = host + ":2289"
	}
	conn, exists := im.Connections.Get(host)
	if exists {
		return conn.(*grpc.ClientConn), nil
	}
	return im.connect(host)
}

// Get gets invite data from a host and invite name
func (im Manager) Get(host, invite string) (InviteData, error) {
	conn, err := im.connectOrCached(host)
	if err != nil {
		return InviteData{}, err
	}

	client := chatv1.NewChatServiceClient(conn)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*15))
	defer cancel()
	data, err := client.PreviewGuild(ctx, &chatv1.PreviewGuildRequest{
		InviteId: invite,
	})

	if err != nil {
		return InviteData{}, err
	}

	avatarHost := ""
	avatarID := ""

	if strings.HasPrefix(data.Avatar, "hmc://") {
		trimmed := strings.TrimPrefix(data.Avatar, "hmc://")
		split := strings.Split(trimmed, "/")
		if len(split) == 2 {
			avatarHost = split[0]
			avatarID = split[1]
		}
	} else {
		avatarHost = host
		avatarID = data.Avatar
	}

	return InviteData{
		GuildName:   data.Name,
		MemberCount: int(data.MemeberCount),
		Avatar:      avatarID,
		AvatarHost:  avatarHost,
		Host:        host,
		InviteID:    invite,
	}, nil
}

func (im Manager) onConnectionEvict(key, value interface{}) {
	conn := value.(*grpc.ClientConn)
	if err := conn.Close(); err != nil {
		log.Println(err)
	}
}
