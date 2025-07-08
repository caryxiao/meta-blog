package main

import (
	"fmt"
	"github.com/caryxiao/meta-blog/db"
	"github.com/caryxiao/meta-blog/internal/config"
	"github.com/caryxiao/meta-blog/internal/di"
	"github.com/caryxiao/meta-blog/internal/middleware"
	"github.com/caryxiao/meta-blog/internal/router"
	"github.com/caryxiao/meta-blog/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// Initialize JWT configuration
	utils.InitJWT(&cfg.JWT)

	var DB *gorm.DB
	if DB, err = db.Init(cfg); err != nil {
		log.Fatal(err)
	}

	c := di.NewContainer(DB)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Tracer())
	router.InitRouter(r, c)
	err = r.Run(fmt.Sprintf(":%s", cfg.App.Port))
	if err != nil {
		return
	}
}
