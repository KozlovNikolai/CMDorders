package main

import (
	"github.com/KozlovNikolai/CMDorders/internal/pkg/config"
	"github.com/KozlovNikolai/CMDorders/internal/server"
)

func main() {
	cfg := config.MustLoad()

	server := server.NewServer(cfg)

	server.Run()
}
