package main

import (
	"log"

	"github.com/korzepadawid/cules-bot/config"
	"github.com/korzepadawid/cules-bot/roll"
	"github.com/korzepadawid/cules-bot/twitter"
	"go.uber.org/zap"
)

const (
	RollCapcity = 4
)

func main() {
	// logger setup
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	// config setup
	c, err := config.Load("./")
	if err != nil {
		log.Fatal(err)
	}
	
	// twitter setup
	tweet := twitter.New(c, logger)
	if tweet.CleanStreamRules(); err != nil {
		log.Fatal(err)
	}

	if err := tweet.CreateStreamRules([]twitter.StreamRule{
		{
			Rule: c.TwitterStreamRule,
			Tag: "sometest",
		},
	}); err != nil {
		log.Fatal(err)
	}

	tweetChann, err := tweet.Stream()
	if err != nil {
		log.Fatal(err)
	}

	r := roll.New[twitter.Tweet](RollCapcity)

	// listener
	go listen(tweetChann, r, logger)
	select{}
}
