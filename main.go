package main

import (
	"log"
	"net/http"
	"text/template"
)

var invitePageTemplate *template.Template

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

	http.HandleFunc("/", handler)
	http.Handle("/static", http.FileServer(http.Dir("./hm-invite/static")))

	log.Fatal(http.ListenAndServe(":2290", nil))
}
