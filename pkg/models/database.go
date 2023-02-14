package models

import "log"

type DatabaseModel struct {
	SessionID string
	Service   string
	Pod       string
	Ingress   string
	Port      string
}

type Database struct {
	Values map[string]DatabaseModel
}

var DB *Database

func init() {
	val := make(map[string]DatabaseModel)
	DB = &Database{val}
}

func AddValuesToDB(ans DatabaseModel) {
	DB.Values[ans.SessionID] = ans
	log.Printf("Written Values to DB %v\n", DB)
}
func GetSessionIDObject(sessionID string) (dbVal DatabaseModel) {

	dbVal = DB.Values[sessionID]
	log.Printf("Got the session:%s object:%v", dbVal)
	return
}
