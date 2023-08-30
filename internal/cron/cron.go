package cron

import (
	"github.com/oskov/megabot/internal/users"
	"log"
	"time"
)

func RunCrons() {
	go runCrons()
}

func runCrons() {
	ticker := time.NewTicker(time.Minute * 10)
	uRep, _ := users.GetUserRepository()
	for {
		select {
		case <-ticker.C:
			err := uRep.IncreaseEnergyGlobal()
			if err != nil {
				log.Println("Error while increasing global energy", err)
			}
		}
	}
}
