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
	
	twitter.New(c, logger).CleanStreamRules()
}
