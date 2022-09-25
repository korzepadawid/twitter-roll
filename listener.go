package main

import (
	"bytes"
	"encoding/json"

	"github.com/korzepadawid/twitter-roll/roll"
	"github.com/korzepadawid/twitter-roll/twitter"
	"go.uber.org/zap"
)

func listen(tweetChann <-chan twitter.Tweet, r *roll.Roll[twitter.Tweet], logger *zap.Logger)  {
	for t := range tweetChann {
		tweetJSONBuff := new(bytes.Buffer)
		json.NewEncoder(tweetJSONBuff).Encode(t)
		logger.Info(tweetJSONBuff.String())
		r.Add(t)
	}
}