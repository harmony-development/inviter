package manager

import (
	"strings"

	chatv1 "github.com/harmony-development/inviter/gen/chat/v1"
	lru "github.com/hashicorp/golang-lru"
)

// Manager manages the connections to various Harmony homeservers
type Manager struct {
	Connections *lru.Cache
}

// New creates a new Manager
func New() *Manager {
	manager := &Manager{}

	cache, err := lru.New(4096)
	if err != nil {
		panic(err)
	}

	manager.Connections = cache
	return manager
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

// Get gets invite data from a host and invite name
func (im Manager) Get(host, invite string) (InviteData, error) {
	client := chatv1.NewChatServiceClient(host)
	data, err := client.PreviewGuild(&chatv1.PreviewGuildRequest{
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
		MemberCount: int(data.MemberCount),
		Avatar:      avatarID,
		AvatarHost:  avatarHost,
		Host:        host,
		InviteID:    invite,
	}, nil
}
