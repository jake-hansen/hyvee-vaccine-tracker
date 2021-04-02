package tweet

import (
	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
)

type Twitter struct {
	Client *twitter.Client
}

func New() *twitter.Client {
	oauthConfig := oauth1.NewConfig("", "")
	token := oauth1.NewToken("", "")
	httpClient := oauthConfig.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
