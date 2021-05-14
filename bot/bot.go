package bot

import (
	"github.com/jake-hansen/hyvee-vaccine-search/adapters"
	"github.com/jake-hansen/hyvee-vaccine-search/api"
	"github.com/jake-hansen/hyvee-vaccine-search/domain"
	"log"
	"time"
)

type Bot struct {
	api        api.HyVeeAPI
	deliverers []domain.Deliverer
	repo 	    *domain.PharmacyMap
	searchParams api.Variables
	logger		*log.Logger
}

type Config struct {
	API 	   api.HyVeeAPI
	Deliverers []domain.Deliverer
	Repo	   *domain.PharmacyMap
	SearchParams api.Variables
	Log			*log.Logger
}

func NewBot(cfg Config) Bot {
	return Bot{
		api:        cfg.API,
		deliverers: cfg.Deliverers,
		repo: cfg.Repo,
		searchParams: cfg.SearchParams,
		logger: cfg.Log,
	}
}

func (b *Bot) StartBot(done chan bool, ticker *time.Ticker) {
	b.updatePharmacies(b.repo)

	for  {
		select {
		case <-ticker.C:
			b.updatePharmacies(b.repo)
		case <- done:
			ticker.Stop()
			return
		}
	}
}

func (b *Bot) updatePharmacies(pharmacyRepo *domain.PharmacyMap) {
	log.Printf("Updating pharmacies... at %s\n", time.Now())

	newPharmaciesStatuses, err := b.getPharmacyMap()
	if err != nil {
		log.Println(err.Error())
	}

	for _, newData := range newPharmaciesStatuses {
		if oldData, ok := (*pharmacyRepo)[domain.PharmacyID(newData.PhoneNumber)]; ok {
			if !oldData.VaccinationsAvailable && newData.VaccinationsAvailable {
				for _, d := range b.deliverers {
					_ = d.Deliver(*oldData)
				}
			}
		}
		(*pharmacyRepo)[newData.ID] = newData
	}
}

func (b *Bot) getPharmacyMap() (domain.PharmacyMap, error) {
	pharmacies, err := b.api.GetPharmacies(b.searchParams)
	returnMap := make(domain.PharmacyMap)

	for _, pharmacy := range pharmacies {
		p := adapters.HyVeePharmacyToDomainPharmacy(pharmacy)
		returnMap[domain.PharmacyID(pharmacy.Location.PhoneNumber)] = &p
	}

	return returnMap, err
}
