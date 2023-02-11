package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Hash map[string]string
type Images struct {
	Hub  []map[string]string `json:"hub"`
	Node struct {
		Chrome  []map[string]string `json:"chrome"`
		Firefox []map[string]string `json:"firefox"`
	} `json:"node"`
}

func GetChromeNodeImage(requestedImage string) (string, error) {
	for _, val := range LoadedImages.Node.Chrome {
		if val[requestedImage] != "" {
			return val[requestedImage], nil
		}
	}
	return "", fmt.Errorf("Requested Image is not present in the list")
}
func GetFirefoxNodeImage(requestedImage string) (string, error) {
	for _, val := range LoadedImages.Node.Firefox {
		if val[requestedImage] != "" {
			return val[requestedImage], nil
		}
	}
	return "", fmt.Errorf("Requested Image is not present in the list")
}
func GetHubImage(requestedImage string) (string, error) {
	for _, val := range LoadedImages.Hub {
		if val[requestedImage] != "" {
			return val[requestedImage], nil
		}
	}
	return "", fmt.Errorf("Requested Image is not present in the list")
}

var LoadedImages Images

func init() {
	jsonFile, err := os.Open("../images.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(byteValue, &LoadedImages)
	if err != nil {
		log.Println(err)
	}
	_ = err
}
