package migrations

import (
	"embed"
	"github.com/oskov/megabot/internal/db"
	"github.com/pressly/goose/v3"
	"log"
)

//go:embed versions/*
var embedMigrations embed.FS

func RunMigrations() error {
	master, err := db.GetDb()
	if err != nil {
		return err
	}
	log.Println("Start migrations")

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Up(master.DB, "versions"); err != nil {
		return err
	}

	log.Println("End migrations")
	return nil
}
