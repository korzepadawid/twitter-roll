package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/korzepadawid/cules-bot/roll"
	"github.com/korzepadawid/cules-bot/twitter"
	"go.uber.org/zap"
)

func Handler(roll *roll.Roll[twitter.Tweet], logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info(fmt.Sprintf("New request from %s", r.RemoteAddr))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(roll.ReadAll())
	}
}