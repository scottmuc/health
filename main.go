package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/scottmuc/health/strava"
)

func main() {
	var accessToken = flag.String("stravaAccessToken", "<required>", "Strava Access Token")
	flag.Parse()

	if *accessToken == "<required>" {
		log.Fatal("stravaAccessToken is required")
	}

	activityName, activityDate := strava.GetLatestActivity(*accessToken)
	fmt.Printf("%s. Occured on %s\n", activityName, activityDate)
}
