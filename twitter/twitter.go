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

	"github.com/korzepadawid/twitter-roll/config"
	"go.uber.org/zap"
)

const (
	twitterStreamRulesAPIEndpoint = "https://api.twitter.com/2/tweets/search/stream/rules"
)

var (
	ErrGettingExistingStreamRules = errors.New("failed to get existing stream rules")
	ErrRemovingExistingStreamRules = errors.New("failed to remove existing stream rules")
	ErrCreatingNewStreamRules = errors.New("failed to create new stream rules")
)

type TwitterClient struct {
	logger *zap.Logger
	bearerToken string
}

type Twitter interface {
	CleanStreamRules() error
	CreateStreamRules(r []StreamRule) error
	Stream() (<- chan Tweet, error)
}

func New(c *config.Config, l *zap.Logger) *TwitterClient {
	return &TwitterClient{
		logger: l,
		bearerToken: c.TwitterBearerToken,
	}
}

type twitterRequest struct {
	url string
	method string
	body io.Reader
}
func (tC *TwitterClient) newAPIRequest(ctx context.Context, r *twitterRequest) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, r.method, r.url, r.body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tC.bearerToken))
	req.Header.Add("content-type","application/json; charset=utf-8")
	
	return req, nil
}

type twitterStreamResponse struct {
	Data []struct {
		ID    string `json:"id"`
	} `json:"data"`
}

func (tC *TwitterClient) getRulesIDs() ([]string, error) {
	tC.logger.Info("Getting existing stream rules")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()
	
	req, err := tC.newAPIRequest(ctx, &twitterRequest{
		url: twitterStreamRulesAPIEndpoint,
		method: http.MethodGet,
	})
	if err != nil {
		return make([]string, 0), err
	}

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
		return make([]string, 0), ErrGettingExistingStreamRules
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
	json.NewEncoder(reqBodyBuff).Encode(body)
	req, err := tC.newAPIRequest(ctx, &twitterRequest{
		method: http.MethodPost,
		url: twitterStreamRulesAPIEndpoint,
		body: reqBodyBuff,
	})
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	
	if res.StatusCode != http.StatusOK {
		return ErrRemovingExistingStreamRules
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

type StreamRule struct {
	// you can use https://developer.twitter.com/apitools/query
	// to generate a new stream filter rule
	Rule string `json:"value"`
	// you can assign a tag-name to
	// your stream rule
	Tag string `json:"tag"`
}

type twitterCreateStreamRulesRequest struct {
	StreamRules []StreamRule `json:"add"`
}

// creates given stream rules, otherwise returns an error
func (tC *TwitterClient) CreateStreamRules(r []StreamRule) error {
	tC.logger.Info("Creating new stream rules")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	body := twitterCreateStreamRulesRequest{
		StreamRules: r,
	}	
	reqBodyBuff := new(bytes.Buffer)
	json.NewEncoder(reqBodyBuff).Encode(body)
	req, err := tC.newAPIRequest(ctx, &twitterRequest{
		method: http.MethodPost,
		url: twitterStreamRulesAPIEndpoint,
		body: reqBodyBuff,
	})
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return ErrCreatingNewStreamRules
	}

	return nil
}