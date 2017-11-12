package strava

import (
	"time"

	gostrava "github.com/strava/go.strava"
)

// ServiceAPI the public interface to the things that I needs
// from the actual strava API
type ServiceAPI interface {
	GetLatestActivity() (string, time.Time)
}

// Service holder of context to connect to strava API
type Service struct {
	client  *gostrava.Client
	service *gostrava.CurrentAthleteService
}

// NewService creates a new service that is ready to talk to the strava API
// using the specified access token
func NewService(accessToken string) ServiceAPI {
	client := gostrava.NewClient(accessToken)
	return &Service{
		client:  client,
		service: gostrava.NewCurrentAthleteService(client),
	}
}

// GetLatestActivity retrieves the name and time of the latest activity
// for the account that holds the access token.
func (s *Service) GetLatestActivity() (string, time.Time) {
	activities, err := s.service.ListActivities().Page(1).PerPage(1).Do()
	if err != nil {
		// TODO this needs a test!
		panic(err)
	}

	// TODO handle the scenario where there are 0 activties
	return activities[0].Name, activities[0].StartDate
}
