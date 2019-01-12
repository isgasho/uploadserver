package common

import "git.jiangshanhudong.com/golang/httpsvr"

type Err struct {
	code int
	err  error
}

var _ httpsvr.APIErr = new(Err)

// Error ...
func (dr *Err) Error() string {
	if dr.err != nil {
		return dr.err.Error()
	}
	return "ok"
}

// Code ...
func (dr *Err) Code() int {
	return dr.code
}

// GenErr ...
func GenErr(code int, err error) *Err {
	return &Err{code: code, err: err}
}

//定义错误码
const (
	//请求成功
	REQUESTESUCCESS	= 22000
	//请求失败
	REQUESTFAILED	= 22001
)


func NewMultiError(errs []error) error {
	message := "Multiple errors occurred:\n"
	for _, err := range errs {
		message += "\t" + err.Error() + "\n"
	}
	return errors.New(message)
}