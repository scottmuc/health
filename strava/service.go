package strava

import (
	"time"

	strava "github.com/strava/go.strava"
)

// GetLatestActivity retrieves the name and time of the latest activity
// for the account that holds the access token.
func GetLatestActivity(accessToken string) (string, time.Time) {
	client := strava.NewClient(accessToken)
	service := strava.NewCurrentAthleteService(client)

	activities, err := service.ListActivities().Page(1).PerPage(1).Do()
	if err != nil {
		// TODO this needs a test!
		panic(err)
	}

	// TODO handle the scenario where there are 0 activties
	return activities[0].Name, activities[0].StartDate
}
