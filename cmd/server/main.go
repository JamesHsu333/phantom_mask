package main

import (
	"fmt"
	"log"

	"github.com/JamesHsu333/kdan/config"
	"github.com/JamesHsu333/kdan/internal/service"
	"github.com/JamesHsu333/kdan/pkg/logger"
	"github.com/JamesHsu333/kdan/pkg/version"
)

const (
	banner = `
____________________________________O/_______
                                    O\
`
)

func main() {
	log.Println("Starting Service...")
	log.Println(version.PrintVersion())

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	l := logger.NewAppLogger(cfg.Logger)
	l.InitLogger()
	fmt.Printf("\n%s%s", cfg.Server.Description, banner)
	l.Infof("Config: %+v", cfg)
	if err := service.NewService(cfg, l).Run(); err != nil {
		l.Fatal(err)
	}
}
