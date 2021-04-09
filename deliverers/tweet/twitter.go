package tweet

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/jake-hansen/hyvee-vaccine-search/domain"
)

type Twitter struct {
	Client *twitter.Client
}

func New() *Twitter {
	oauthConfig := oauth1.NewConfig("", "")
	token := oauth1.NewToken("", "")
	httpClient := oauthConfig.Client(oauth1.NoContext, token)

	returnTwitter := &Twitter{Client: twitter.NewClient(httpClient)}

	return returnTwitter
}

func (t *Twitter) Deliver(pharmacy domain.Pharmacy) error {
	_, res, err := t.Client.Statuses.Update(pharmacyToTweet(pharmacy), nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if res != nil {
			fmt.Printf("Tweet for %s response code %d", pharmacy.PhoneNumber, res.StatusCode)
		}
	}

	return nil
}

func pharmacyToTweet(pharmacy domain.Pharmacy) string {
	addressLineCombination := pharmacy.Address.Line1
	if pharmacy.Address.Line2 != "" {
		addressLineCombination = addressLineCombination + "\n" + pharmacy.Address.Line2
	}

	url := "https://www.hy-vee.com/my-pharmacy/covid-vaccine-consent"

	return fmt.Sprintf("New appointments available at\n%s\n%s, %s %d\n\nPhone: %s\n\n%s",
		addressLineCombination,
		pharmacy.Address.City,
		pharmacy.Address.State,
		pharmacy.Address.Zip,
		pharmacy.PhoneNumber,
		url)
}

