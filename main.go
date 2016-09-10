package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/scottmuc/health/dates"
	"github.com/scottmuc/health/strava"
)

func main() {
	var accessToken = flag.String("stravaAccessToken", "<required>", "Strava Access Token")
	flag.Parse()

	if *accessToken == "<required>" {
		log.Fatal("stravaAccessToken is required")
	}

	stravaService := strava.NewService(*accessToken)
	activityName, activityDate := stravaService.GetLatestActivity()
	daysAgo := dates.DayDifference(activityDate, time.Now())
	fmt.Printf("%s. Occurred %d day(s) ago\n", activityName, daysAgo)
}
