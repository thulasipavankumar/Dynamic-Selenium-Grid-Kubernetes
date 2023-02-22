package misc

import (
	"os"
	"strconv"

	"github.com/jasonlvhit/gocron"
)

const (
	CLEANUP_MINUTES_INTERVAL = 3
)

var cleanupIntervalMinutes int

func init() {
	cleanupIntervalMinutes, _ = strconv.Atoi(os.Getenv("cleanupIntervalMinutes"))
	if cleanupIntervalMinutes == 0 {
		cleanupIntervalMinutes = CLEANUP_MINUTES_INTERVAL
	}
}
func ExecuteCronJob() {
	gocron.Every(uint64(cleanupIntervalMinutes)).Minutes().Do(ClearStaleSessions)
	<-gocron.Start()
}
