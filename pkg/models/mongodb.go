package models

import "log"

type model struct {
	sessionId   string
	podName     string
	ServiceName string
	ingressName string
}

func WriteCreateSessionToDb(m model) {
	log.Println("Writing to db", m)
}
func updateDeleteSessionToDb() {

}
