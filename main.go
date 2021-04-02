package main

import (
	"fmt"
	"github.com/jake-hansen/hyvee-vaccine-search/api"
	"github.com/jake-hansen/hyvee-vaccine-search/tweet"
	"net/http"
	"time"
)

func main() {
	hyveeAPI := &api.HyVeeAPI{Client: http.DefaultClient}
	
	searchParams := api.Variables{
		Radius:    75,
		Latitude:  41.2354329,
		Longitude: -95.99383390000001,
	}

	foundPharmaciesVaccineAvailable := make(map[string]bool)

	t := tweet.New()

	for true {
		pharmacies := hyveeAPI.GetPharmacies(searchParams)
		newlyAvailable := 0

		fmt.Println("Retrieved at " + time.Now().String())
		for _, pharmacy := range pharmacies {
			availabilityString := "AVAILABLE"
			if !pharmacy.Location.IsCovidVaccineAvailable {
				availabilityString = "NOT AVAILABLE"
			}
			fmt.Printf("%s: %v\n", availabilityString, pharmacy)

			if p, ok := foundPharmaciesVaccineAvailable[pharmacy.Location.PhoneNumber]; ok {
				// vaccine status updated, send email
				if p == false && pharmacy.Location.IsCovidVaccineAvailable {
					newlyAvailable ++
					_, res, err := t.Statuses.Update(pharmacyToTweet(pharmacy), nil)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						if res != nil {
							fmt.Printf("Tweet for %s response code %d", pharmacy.Location.PhoneNumber, res.StatusCode)
						}
					}
				}
			}
			foundPharmaciesVaccineAvailable[pharmacy.Location.PhoneNumber] = pharmacy.Location.IsCovidVaccineAvailable

		}
		fmt.Printf("\nNewly Available: %d\n\n", newlyAvailable)
		newlyAvailable = 0

		time.Sleep(time.Minute)
	}

}

func pharmacyToTweet(pharmacy api.Pharmacy) string {
	addressLineCombination := pharmacy.Location.Address.Line1
	if pharmacy.Location.Address.Line2 != "" {
		addressLineCombination = addressLineCombination + "\n" + pharmacy.Location.Address.Line2
	}


	url := "https://www.hy-vee.com/my-pharmacy/covid-vaccine-consent"

	return fmt.Sprintf("New appointments available at\n%s\n%s, %s %s\n\nPhone: %s\n\n%s",
		addressLineCombination,
		pharmacy.Location.Address.City,
		pharmacy.Location.Address.State,
		pharmacy.Location.Address.Zip,
		pharmacy.Location.PhoneNumber,
		url)
}

func pharmacyToMailMessage(pharmacy api.Pharmacy) string {
	addressLineCombination := pharmacy.Location.Address.Line1
	if pharmacy.Location.Address.Line2 != "" {
		addressLineCombination = addressLineCombination + "\n" + pharmacy.Location.Address.Line2
	}

	url := "https://www.hy-vee.com/my-pharmacy/covid-vaccine-consent"

	return fmt.Sprintf("<body>Name: %s\n<br>Location:\n<br>%s\n<br>%s, %s %s\n<br>Phone Number: %s\n<br><br>\n <a href=\"%s\">%s</a></body>",
		pharmacy.Location.Nickname,
		addressLineCombination,
		pharmacy.Location.Address.City,
		pharmacy.Location.Address.State,
		pharmacy.Location.Address.Zip,
		pharmacy.Location.PhoneNumber,
		url,
		url)
}





