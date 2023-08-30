package main

import (
	"github.com/joho/godotenv"
	"github.com/oskov/megabot/internal/bot"
	"github.com/oskov/megabot/internal/config"
	"github.com/oskov/megabot/internal/cron"
	"github.com/oskov/megabot/migrations"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	rand.Seed(time.Now().Unix())
}

func main() {
	cnfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	cron.RunCrons()

	err = migrations.RunMigrations()
	if err != nil {
		panic(err)
	}

	b, err := bot.GetBot(cnfg)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = <-c
		log.Println("Stop bot")
		b.Stop()
	}()

	log.Println("Start bot")
	b.Start()

	log.Println("Exit")
}
