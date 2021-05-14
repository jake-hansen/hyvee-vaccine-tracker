package main

import (
	"github.com/jake-hansen/hyvee-vaccine-search/api"
	"github.com/jake-hansen/hyvee-vaccine-search/bot"
	"github.com/jake-hansen/hyvee-vaccine-search/deliverers/consoleprinter"
	"github.com/jake-hansen/hyvee-vaccine-search/deliverers/tweet"
	"github.com/jake-hansen/hyvee-vaccine-search/domain"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	searchParams := api.Variables{
		Radius:    75,
		Latitude:  41.2354329,
		Longitude: -95.99383390000001,
	}
	
	twitterConfig := tweet.DefaultConfig(
		os.Getenv("COVID_TOKEN"),
		os.Getenv("COVID_TOKEN_SECRET"),
		os.Getenv("COVID_CONSUMER_KEY"),
		os.Getenv("COVID_CONSUMER_SECRET"))

	logger := log.Default()
	pharmacyRepo := make(domain.PharmacyMap)
	deliverers := []domain.Deliverer{tweet.New(twitterConfig) ,consoleprinter.New()}
	
	botConfig := bot.Config{
		API:          api.HyVeeAPI{Client: http.DefaultClient},
		Deliverers:   deliverers,
		Repo:         &pharmacyRepo,
		SearchParams: searchParams,
		Log: logger,
	}

	done := make(chan bool)
	ticker := time.NewTicker(time.Minute)

	newBot := bot.NewBot(botConfig)
	newBot.StartBot(done, ticker)
}
