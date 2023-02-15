package utils

import (
	"fmt"
	"log"
)

type Response struct {
	ResData      []byte
	Err          error
	ResponseCode int
	M            interface{}
}

func (r Response) GetResponseData() (byteSlice []byte) {
	return r.ResData
}
func (r Response) GetErr() error {
	return r.Err
}
func (r Response) GetResponseCode() int {
	return r.ResponseCode
}
func (r Response) Printf(m interface{}) {
	log.Printf("%s data: %s error:%v responcecode:%d \n", m, fmt.Sprintf("%s", r.ResData), r.Err, r.ResponseCode)
}
func (r Response) Println(m interface{}) {
	log.Printf("%s data: %s error:%v responcecode:%d \n", m, fmt.Sprintf("%s", r.ResData), r.Err, r.ResponseCode)
}
