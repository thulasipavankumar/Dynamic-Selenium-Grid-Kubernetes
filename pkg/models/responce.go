package models

type Response struct {
	resData      []byte
	err          error
	responseCode int
}

func (r Response) GetResponseData() (byteSlice []byte) {
	return r.resData
}
func (r Response) GetErr() error {
	return r.err
}
func (r Response) GetResponseCode() int {
	return r.responseCode
}
