package main

import (
	"github.com/harmony-development/inviter/app"
	"github.com/harmony-development/inviter/app/manager"
)

func main() {
	app.Start(manager.New())
}
