package main

import (
	"github.com/KozlovNikolai/CMDorders/internal/config"
	"github.com/KozlovNikolai/CMDorders/internal/server"
)

func main() {
	cfg := config.MustLoad()

	server := server.NewServer(cfg)

	server.Run()
}
