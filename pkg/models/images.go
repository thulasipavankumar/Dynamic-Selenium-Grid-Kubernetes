package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (images Images) getNodeImage(requestedImage string) (string, error) {
	for _, val := range images.Node.Chrome {
		if val[requestedImage] != "" {
			return val[requestedImage], nil
		}
	}
	return "", fmt.Errorf("Requested Image is not present in the list")
}

var LoadedImages Images

func init() {
	jsonFile, err := os.Open("../pkg/models/images.json")
	byteValue, err := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &LoadedImages)
	_ = err
}
