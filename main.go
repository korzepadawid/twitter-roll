package main

import (
	"log"

	"github.com/korzepadawid/cules-bot/config"
	"github.com/korzepadawid/cules-bot/twitter"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	c, err := config.Load("./")
	if err != nil {
		log.Fatal(err)
	}
	
	tweet := twitter.New(c, logger)
	if tweet.CleanStreamRules(); err != nil {
		log.Fatal(err)
	}
}
