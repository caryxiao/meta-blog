package main

import (
	"github.com/caryxiao/meta-blog/internal/config"
	"log"
	"os"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	cfg, err := config.LoadConfig(env)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)
}
