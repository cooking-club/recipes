package main

import (
	"github.com/cooking-club/recipes/internal/db"
	"github.com/cooking-club/recipes/internal/server"
)

func main() {
	db.Init()
	defer db.Close()
	server.Run()
}
