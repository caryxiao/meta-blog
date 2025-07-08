package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/caryxiao/meta-blog/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"path/filepath"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	var action string
	flag.StringVar(&action, "action", "up", "Migration action: up / down / force / drop")
	flag.Parse()

	cfg, err := config.LoadConfig(env)

	if err != nil {
		log.Fatalf("Loading App Config Failed: %v", err)
	}

	migrationsFilePath := "file://" + filepath.Join("db", "migrations")
	m, err := migrate.New(migrationsFilePath, cfg.Database.MigrateURL())
	if err != nil {
		log.Fatalf("Initialize Migrate Failed: %v", err)
	}

	switch action {
	case "up":
		err = m.Up()
	case "down":
		err = m.Steps(-1)
	case "drop":
		err = m.Drop()
	case "force":
		err = m.Force(1)
	default:
		log.Fatalf("Not support action: %s", action)
	}

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Excution Failed: %v", err)
	}

	fmt.Println("Database Migration Complete")
}
