package misc

import (
	"encoding/json"
	"log"
	"time"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
	kubernetes "github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models/kubernetes"
	selenium "github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models/selenium"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/utils"
)

func getUnclosedSessionMoreThan3Minuts() []models.DatabaseModel {
	sessions := models.GetAllOpenSessions()
	var staleSessions []models.DatabaseModel
	for _, session := range sessions {

		diff := time.Now().Sub(session.CreatedAt)
		hours, minutes, seconds := diff.Hours(), diff.Minutes(), diff.Seconds()
		if minutes > 3 {
			//fmt.Println("Succes found entry > 3 minutes ", session)
			staleSessions = append(staleSessions, session)
		}
		_, _ = hours, seconds

	}
	return staleSessions
}

func ClearStaleSessions() {
	staleSessions := getUnclosedSessionMoreThan3Minuts()
	payload := struct {
		Query string `json:"query"`
	}{
		Query: "{ grid { uri, maxSession, sessionCount }, nodesInfo { nodes { id, uri, status, sessions { id, capabilities, startTime, uri, nodeId, nodeUri, sessionDurationMillis, slot { id, stereotype, lastStarted } }, slotCount, sessionCount }} }",
	}

	for _, session := range staleSessions {
		_ = session
		bytes, err := json.Marshal(payload)
		if err != nil {
			log.Println("Error in getting sessions from the hub ", err)
			break
		}

		responce := utils.Make_Post_Call("http://"+session.Service+":4444/graphql", bytes)
		var gridSessions selenium.GridSessions
		unmarshalError := json.Unmarshal([]byte(responce.ResData), &gridSessions)
		//responce.Printf("Session responce is ")
		if gridSessions.Data.Grid.SessionCount == 0 && responce.ResponseCode == 200 {

			log.Printf("session can be deleted %v \n", session)
			deployment := kubernetes.Deployment{}
			go deployment.DeleteDeployment(session.Pod, session.Service, session.Ingress)
			go models.DeleteDBCell(session.SessionID)
		}
		_ = unmarshalError

	}
}
