package main

import (
	"log"
	"net/http"
	"text/template"

	chatv1 "github.com/harmony-development/inviter/gen/chat/v1"
)

var invitePageTemplate *template.Template
var client *chatv1.ChatServiceClient

// InviteData contains the data for a guild
type InviteData struct {
	GuildName   string
	MemberCount int
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := InviteData{GuildName: "Hello"}
	invitePageTemplate.Execute(w, data)
}

func main() {
	parsed, err := template.ParseFiles("./hm-invite/index.html")
	if err != nil {
		panic(err)
	}

	invitePageTemplate = parsed

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./hm-invite/static"))))
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":2290", nil))
}
