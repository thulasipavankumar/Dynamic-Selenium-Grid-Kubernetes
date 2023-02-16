package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type MemoryResource struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}
type Resources struct {
	Limits   MemoryResource `json:"limits"`
	Requests MemoryResource `json:"requests"`
}
type Hash map[string]string
type images struct {
	Hub  []map[string]string `json:"hub"`
	Node struct {
		Chrome  []map[string]string `json:"chrome"`
		Firefox []map[string]string `json:"firefox"`
	} `json:"node"`
}
type imageresources struct {
	Hub []struct {
		ResourcesObj Resources `json:"resources"`
	} `json:"hub"`
	Node struct {
		Chrome []struct {
			Str          map[string]string `json:"omitempty,omitempty"`
			ResourcesObj Resources         `json:"resources"`
		} `json:"chrome"`
		Firefox []struct {
			Str          map[string]string `json:","`
			ResourcesObj Resources         `json:"resources"`
		} `json:"firefox"`
	} `json:"node"`
}

func GetChromeNodeImage(requestedImage string) (Image, error) {

	if chromeImages[requestedImage].ImageValue != "" {
		return chromeImages[requestedImage], nil
	}
	return Image{}, fmt.Errorf("Requested Image is not present in the list")
}
func GetFirefoxNodeImage(requestedImage string) (Image, error) {

	if firefoxImages[requestedImage].ImageValue != "" {
		return firefoxImages[requestedImage], nil
	}
	return Image{}, fmt.Errorf("Requested Image is not present in the list")
}
func GetHubImage(requestedImage string) (Image, error) {

	if hubImages[requestedImage].ImageValue != "" {
		return hubImages[requestedImage], nil
	}

	return Image{}, fmt.Errorf("Requested Image is not present in the list")
}

type Image struct {
	ImageValue  string
	ResourceObj Resources
}

var hubImages map[string]Image
var chromeImages map[string]Image
var firefoxImages map[string]Image
var imagesobj images
var resourcesObj imageresources

func init() {
	hubImages = make(map[string]Image)
	chromeImages = make(map[string]Image)
	firefoxImages = make(map[string]Image)
	jsonFile, err := os.Open("../images.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(byteValue, &imagesobj)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(byteValue, &resourcesObj)
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < len(resourcesObj.Hub); i++ {
		for key, val := range imagesobj.Hub[i] {
			hubImages[key] = Image{ImageValue: val, ResourceObj: resourcesObj.Hub[i].ResourcesObj}
			break
		}

	}
	for i := 0; i < len(resourcesObj.Node.Chrome); i++ {

		for key, val := range imagesobj.Node.Chrome[i] {
			chromeImages[key] = Image{ImageValue: val, ResourceObj: resourcesObj.Node.Chrome[i].ResourcesObj}
			break
		}

	}
	for i := 0; i < len(resourcesObj.Node.Firefox); i++ {
		for key, val := range imagesobj.Node.Firefox[i] {
			firefoxImages[key] = Image{ImageValue: val, ResourceObj: resourcesObj.Node.Firefox[i].ResourcesObj}
			break
		}

	}
	fmt.Printf("hub image :%v\n", hubImages)
	fmt.Printf("chrome image :%v\n", chromeImages)
	fmt.Printf("firefox image :%v\n", firefoxImages)
	_ = err
}
