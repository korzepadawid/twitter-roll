package twitter

import "github.com/korzepadawid/cules-bot/config"

type TwitterClient struct {
	bearerToken string
}

type Twitter interface {
	CleanFilters()
}

func New(c *config.Config) *TwitterClient {
	return &TwitterClient{
		bearerToken: c.TwitterBearerToken,
	}
}