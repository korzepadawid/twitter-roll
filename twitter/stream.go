package twitter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	twitterStreamAPIEndpoint = "https://api.twitter.com/2/tweets/search/stream?tweet.fields=attachments,author_id,id,referenced_tweets,text"
)

type Tweet struct {
	Data struct {
		ID string `json:"id"`
		AuthorID string `json:"author_id"`
		Text string `json:"text"`
	} `json:"data"`
}

func (tC *TwitterClient) Stream() (<- chan Tweet, error) {
	tweetChann := make(chan Tweet)
	req, err := http.NewRequest(http.MethodGet, twitterStreamAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tC.bearerToken))
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	
	tC.logger.Info("Starting to listen for tweets")	
	go func (body io.Reader)  {
		for {
			var tweet Tweet
			json.NewDecoder(body).Decode(&tweet)
			tweetChann <- tweet
		}
	}(res.Body)
	
	return tweetChann, nil
}