package config

import (
	"os"
	"strconv"
)

var port int
var service string

func init() {
	port, _ = strconv.Atoi(os.Getenv("Port"))
	if port == 0 {
		port = 8080
	}
	service = os.Getenv("Service")
}
func GetPort() int {
	return port
}
func GetService() string {
	return service
}
