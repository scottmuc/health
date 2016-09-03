package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/strava/go.strava"
)

func main() {
	var accessToken = flag.String("stravaAccessToken", "<required>", "Strava Access Token")
	flag.Parse()

	if *accessToken == "<required>" {
		fmt.Println("stravaAccessToken is required")
		os.Exit(1)
	}

	client := strava.NewClient(*accessToken)
	service := strava.NewCurrentAthleteService(client)

	activities, err := service.ListActivities().Page(1).PerPage(1).Do()
	if err != nil {
		panic(err)
	}

	latestActivity := activities[0]
	fmt.Println(latestActivity)
}
