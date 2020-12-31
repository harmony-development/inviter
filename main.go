package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	chatv1 "github.com/harmony-development/inviter/gen/chat/v1"
)

var invitePageTemplate *template.Template
var connMgr *Manager

// InviteData contains the data for a guild
type InviteData struct {
	GuildName   string
	MemberCount int
	Host        string
	InviteID    string
}

func handler(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimPrefix(r.URL.Path, "/")
	split := strings.Split(trimmed, "/")

	if len(split) < 2 {
		w.Write([]byte("invalid invite"))
		return
	}

	host := split[0]
	invite := split[1]

	conn, err := connMgr.GetOrConnect(host)
	if err != nil {
		w.Write([]byte("unable to connect to server"))
		log.Println(err)
		return
	}
	c := chatv1.NewChatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	resp, err := c.PreviewGuild(ctx, &chatv1.PreviewGuildRequest{
		InviteId: invite,
	})
	if err != nil {
		w.Write([]byte("unable to load invite"))
		log.Println(err)
		return
	}

	data := InviteData{GuildName: resp.Name, MemberCount: int(resp.MemeberCount), Host: host, InviteID: invite}
	invitePageTemplate.Execute(w, data)
}

func main() {
	parsed, err := template.ParseFiles("./hm-invite/index.html")
	if err != nil {
		panic(err)
	}

	invitePageTemplate = parsed

	mgr, err := NewManager()
	if err != nil {
		panic(err)
	}
	connMgr = mgr

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./hm-invite/static"))))
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":2290", nil))
}
