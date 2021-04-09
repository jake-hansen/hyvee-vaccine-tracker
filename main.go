package main

import (
	"fmt"
	"github.com/jake-hansen/hyvee-vaccine-search/adapters"
	"github.com/jake-hansen/hyvee-vaccine-search/api"
	"github.com/jake-hansen/hyvee-vaccine-search/deliverers/consoleprinter"
	"github.com/jake-hansen/hyvee-vaccine-search/deliverers/tweet"
	"github.com/jake-hansen/hyvee-vaccine-search/domain"
	"net/http"
	"time"
)

type PharmacyMap map[domain.PharmacyID]*domain.Pharmacy

func main() {
	pharmacyRepo := make(PharmacyMap)
	done := make(chan bool)
	ticker := time.NewTicker(time.Minute)
	updatePharmacies(&pharmacyRepo)
	startBot(&pharmacyRepo, done, ticker)
}

type Deliverer interface {
	Deliver(pharmacy domain.Pharmacy) error
}

type Bot struct {
	API       api.HyVeeAPI
	Deliverers []Deliverer
}

func startBot(pharmacyRepo *PharmacyMap, done chan bool, ticker *time.Ticker) {
	for  {
		select {
			case <-ticker.C:
				updatePharmacies(pharmacyRepo)
			case <- done:
				ticker.Stop()
				return
		}
	}
}

func updatePharmacies(pharmacyRepo *PharmacyMap) {
	fmt.Printf("Updating pharmacies... at %s\n", time.Now())
	searchParams := api.Variables{
		Radius:    75,
		Latitude:  41.2354329,
		Longitude: -95.99383390000001,
	}

	deliverers := []Deliverer{tweet.New() ,consoleprinter.New()}

	bot := Bot{
		API:       api.HyVeeAPI{Client: http.DefaultClient},
		Deliverers: deliverers,
	}

	newPharmaciesStatuses := getPharmacyMap(bot.API, searchParams)

	for _, pharmacy := range newPharmaciesStatuses {
		if p, ok := (*pharmacyRepo)[domain.PharmacyID(pharmacy.PhoneNumber)]; ok {
			if p.VaccinationsAvailable == false && pharmacy.VaccinationsAvailable {
				for _, d := range bot.Deliverers {
					_ = d.Deliver(*p)
				}
			}
		}
		(*pharmacyRepo)[pharmacy.ID] = pharmacy
	}
}

func getPharmacyMap(api api.HyVeeAPI, params api.Variables) PharmacyMap {
	pharmacies := api.GetPharmacies(params)
	returnMap := make(PharmacyMap)

	for _, pharmacy := range pharmacies {
		p := adapters.HyVeePharmacyToDomainPharmacy(pharmacy)
		returnMap[domain.PharmacyID(pharmacy.Location.PhoneNumber)] = &p
	}

	return returnMap
}
