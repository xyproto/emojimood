package emojigo

import (
	"fmt"
	api "github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

func Connect() (*Slack, error) {
	s := Slack{}

	s.Api = api.New(Config.SlackToken)
	auth, err := s.Api.AuthTest()
	if err != nil {
		return &s, fmt.Errorf("Error authenticating with Slack: %s", err)
	}

	log.WithFields(log.Fields{
		"teamName": auth.Team,
		"userId":   auth.UserID,
		"teamUrl":  auth.URL,
	}).Info("Authenticated with Slack")

	return &s, nil
}
