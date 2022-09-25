package main

import (
	"log"
	"net/http"

	"github.com/korzepadawid/cules-bot/config"
	"github.com/korzepadawid/cules-bot/handler"
	"github.com/korzepadawid/cules-bot/roll"
	"github.com/korzepadawid/cules-bot/twitter"
	"go.uber.org/zap"
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
			Tag: "twitter-bot",
		},
	}); err != nil {
		log.Fatal(err)
	}

	tweetChann, err := tweet.Stream()
	if err != nil {
		log.Fatal(err)
	}

	r := roll.New[twitter.Tweet](c.RollCapcity)

	// listener
	go listen(tweetChann, r, logger)

	// http server
	logger.Info("Starting HTTP server")
	http.HandleFunc("/roll", handler.Handler(r, logger))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
