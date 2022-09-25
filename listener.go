package main

import (
	"fmt"

	"github.com/korzepadawid/cules-bot/roll"
	"github.com/korzepadawid/cules-bot/twitter"
	"go.uber.org/zap"
)

func listen(tweetChann <-chan twitter.Tweet, r *roll.Roll[twitter.Tweet], logger *zap.Logger)  {
	for t := range tweetChann {
		tweetDetails := fmt.Sprintf("TweetID=%s Text=%s AuthorID=%s", t.Data.ID, t.Data.Text, t.Data.AuthorID)
		logger.Info(tweetDetails)
		r.Add(t)
	}
}