package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/scottmuc/health/dates"
	"github.com/scottmuc/health/strava"
)

func main() {
	var accessToken = flag.String("stravaAccessToken", "", "Strava Access Token")
	flag.Parse()

	if *accessToken != "" {
		stravaService := strava.NewService(*accessToken)
		activityName, activityDate := stravaService.GetLatestActivity()
		daysAgo := dates.DayDifference(activityDate, time.Now())
		fmt.Printf("%s. Occurred %d day(s) ago\n", activityName, daysAgo)
	} else {
		fmt.Printf("No stravaAccessToken provided, ignoring...")
	}
}
