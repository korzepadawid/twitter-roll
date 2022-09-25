package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	twitterStreamAPIEndpoint = "https://api.twitter.com/2/tweets/search/stream?tweet.fields=attachments,author_id,created_at,text&expansions=author_id&media.fields=preview_image_url,url&user.fields=profile_image_url,url,username"
)

var (
	ErrTwitterStream = errors.New("failed to stream tweets")
)

type Tweet struct {
	Data          Data            `json:"data"`
	Includes      Includes        `json:"includes"`
}

type Data struct {
	AuthorID  string    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	Text      string    `json:"text"`
}

type Users struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url"`
	URL             string `json:"url"`
	Username        string `json:"username"`
}

type Includes struct {
	Users []Users `json:"users"`
}


func (tC *TwitterClient) Stream() (<- chan Tweet, error) {
	req, err := http.NewRequest(http.MethodGet, twitterStreamAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tC.bearerToken))
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	
	if res.StatusCode != http.StatusOK {
		return nil, ErrTwitterStream
	}

	tC.logger.Info("Starting to listen for tweets")	
	tweetChann := make(chan Tweet)
	go func (body io.Reader)  {
		for {
			var tweet Tweet
			if err := json.NewDecoder(body).Decode(&tweet); err == nil {
				tweetChann <- tweet
			}
		}
	}(res.Body)
	
	return tweetChann, nil
}