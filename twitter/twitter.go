package twitter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/korzepadawid/cules-bot/config"
	"go.uber.org/zap"
)

var (
	ErrCantGetExistingStreamRules = errors.New("failed to get existing stream rules")
)

type TwitterClient struct {
	logger *zap.Logger
	bearerToken string
}

type Twitter interface {
	CleanStreamRules() error
}

func New(c *config.Config, l *zap.Logger) *TwitterClient {
	return &TwitterClient{
		logger: l,
		bearerToken: c.TwitterBearerToken,
	}
}

type twitterFiltersResponse struct {
	Data []struct {
		ID    string `json:"id,omitempty"`
	} `json:"data"`
}

func (tC *TwitterClient) getRulesIDs() (*twitterFiltersResponse, error) {
	tC.logger.Info("Getting existing stream rules")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.twitter.com/2/tweets/search/stream/rules", nil)
	if err != nil {
		return &twitterFiltersResponse{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tC.bearerToken))
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &twitterFiltersResponse{} , err
	}
	defer func(b io.ReadCloser) {
		err := b.Close()
		if err != nil {
			return
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return &twitterFiltersResponse{}, ErrCantGetExistingStreamRules
	}

	var data twitterFiltersResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil{
		return &twitterFiltersResponse{}, err
	}

	return &data, nil
}

// Removes already existing stream rules for your Twitter account.
// Results won't be affected by unwanted rules.
func (tC *TwitterClient) CleanStreamRules() error {
	rules, err := tC.getRulesIDs()
	if err != nil {
		return nil
	}
	rulesIDs := rules.Data
	fmt.Println(rulesIDs)
	return nil
}