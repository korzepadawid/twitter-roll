package main

import (
	"log"
	"net/http"

	"github.com/korzepadawid/twitter-roll/config"
	"github.com/korzepadawid/twitter-roll/handler"
	"github.com/korzepadawid/twitter-roll/roll"
	"github.com/korzepadawid/twitter-roll/twitter"
	"github.com/rs/cors"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/roll", handler.Handler(r, logger))
	muxHandler := cors.Default().Handler(mux)
	if err := http.ListenAndServe(":8080", muxHandler); err != nil {
		log.Fatal(err)
	}
}
