package utils

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
