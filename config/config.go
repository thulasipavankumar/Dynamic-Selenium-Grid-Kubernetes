package config

import (
	"os"
	"strconv"
)

var Port int
var Service string

func init() {
	Port, _ = strconv.Atoi(os.Getenv("Port"))
	if Port == 0 {
		Port = 8080
	}
	Service = os.Getenv("Service")

}
