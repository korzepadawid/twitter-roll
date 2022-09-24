package twitter

import (
	"bytes"
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
	ErrFailedToGetExistingStreamRules = errors.New("failed to get existing stream rules")
	ErrFailedToRemoveExistingStreamRules = errors.New("failed to remove existing stream rules")
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

type twitterStreamResponse struct {
	Data []struct {
		ID    string `json:"id,omitempty"`
	} `json:"data"`
}

func (tC *TwitterClient) getRulesIDs() ([]string, error) {
	tC.logger.Info("Getting existing stream rules")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.twitter.com/2/tweets/search/stream/rules", nil)
	if err != nil {
		return make([]string, 0), err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tC.bearerToken))
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return make([]string, 0) , err
	}
	defer func(b io.ReadCloser) {
		err := b.Close()
		if err != nil {
			return
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return make([]string, 0), ErrFailedToGetExistingStreamRules
	}

	var data twitterStreamResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil{
		return make([]string, 0), err
	}

	ids := make([]string, len(data.Data))
	for idx, d := range data.Data {
		ids[idx] = d.ID
	}

	return ids, nil
}

type twitterCleanRulesRequest struct {
	Delete delete `json:"delete"`
}

type delete struct {
	Ids []string `json:"ids"`
}

func (tC *TwitterClient) deleteStreamRules(ids []string) error {
	tC.logger.Info("Deleting stream rules")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	body := twitterCleanRulesRequest{
		Delete: delete{
			Ids: ids,
		},
	}
	reqBodyBuff := new(bytes.Buffer)
	json.NewEncoder(reqBodyBuff).Encode(&body)
	req, err := http.NewRequestWithContext(ctx,http.MethodPost, "https://api.twitter.com/2/tweets/search/stream/rules", reqBodyBuff)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tC.bearerToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return ErrFailedToRemoveExistingStreamRules
	}

	return nil
}

// Removes already existing stream rules for your Twitter account.
// Results won't be affected by unwanted rules.
func (tC *TwitterClient) CleanStreamRules() error {
	rules, err := tC.getRulesIDs()
	if err != nil {
		return err
	}

	if err := tC.deleteStreamRules(rules); err != nil {
		return err
	}

	return nil
}