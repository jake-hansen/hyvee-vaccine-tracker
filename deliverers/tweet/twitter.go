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

type Config struct {
	OAuthConfig *oauth1.Config
	Token oauth1.Token
}

func DefaultConfig(token string, tokenSecret string, consumerKey string, consumerSecret string) Config {
	return Config{
		OAuthConfig: oauth1.NewConfig(consumerKey, consumerSecret),
		Token:       oauth1.Token{
			Token:       token,
			TokenSecret: tokenSecret,
		},
	}
}

func New(cfg Config) *Twitter {
	httpClient := cfg.OAuthConfig.Client(oauth1.NoContext, &cfg.Token)

	return &Twitter{Client: twitter.NewClient(httpClient)}
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

