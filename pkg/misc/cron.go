package misc

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
)

func getUnclosedSessionMoreThan3Minuts() {
	sessions := models.GetAllOpenSessions()

	for _, session := range sessions {

		diff := time.Now().Sub(session.CreatedAt)
		hours, minutes, seconds := diff.Hours(), diff.Minutes(), diff.Seconds()
		if minutes > 3 {
			fmt.Println("Succes found entry > 3 minutes ", session)

		}
		_, _ = hours, seconds

	}
}
func ExecuteCronJob() {
	gocron.Every(3).Minutes().Do(getUnclosedSessionMoreThan3Minuts)
	<-gocron.Start()
}
