package main

import (
	"os"

	"github.com/hu553in/gist-clone/config"
	"github.com/hu553in/gist-clone/db"
	"github.com/hu553in/gist-clone/server"
)

func main() {
	env := os.Getenv("GIST_CLONE_ENV")
	if env == "" {
		env = "dev"
	}

	config.Init(env)
	db.Init()
	server.Init()
}
