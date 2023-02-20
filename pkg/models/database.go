package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/config"
)

type DatabaseModel struct {
	gorm.Model
	SessionID      string `gorm:"primaryKey"`
	Service        string
	Pod            string
	Ingress        string
	SessionUrl     string
	Browser        string
	BrowserVersion string
}

// type Database struct {
// 	Values map[string]DatabaseModel
// }

//var DB *Database

func init() {
	// val := make(map[string]DatabaseModel)

}

func AddValuesToDB(ans DatabaseModel) {
	//DB.Values[ans.SessionID] = ans

	val := config.GetDB().Create(&ans)
	log.Printf("Written Values to DB %v \n %v\n", ans, val)
}
func GetSessionIDObject(sessionID string) (dbVal DatabaseModel) {

	// dbVal = DB.Values[sessionID]
	// log.Printf("Got the session:%s object:%v", sessionID, dbVal)
	config.GetDB().Where("session_id =?", sessionID).Find(&dbVal)
	log.Printf("Got the session:%s object:%v", sessionID, dbVal)
	return
}
func DeleteDBCell(sessionID string) {
	var dbRow DatabaseModel
	config.GetDB().Where("session_id =?", sessionID).Delete(dbRow)
	log.Printf("Delete db row for %s : %v ", sessionID, dbRow)
}
func GetAllOpenSessions() (sessions []DatabaseModel) {
	config.GetDB().Where("deleted_at IS NULL").Find(&sessions)
	return
}
