package utils

//状态码
const (
	stSucc    int = 200 //正常
	stFail    int = 300 //失败
	stErrIpt  int = 310 //输入数据有误
	stErrOpt  int = 320 //无数据返回
	stErrDeny int = 330 //没有权限
	stErrJwt  int = 340 //jwt未通过验证
	stErrSvr  int = 350 //服务端错误
	stExt     int = 400 //其他约定 //eg 更新 token
)

// 响应数据
type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

//页码
type page struct {
	Count int         `json:"count"`
	Items interface{} `json:"items"`
}

func newReply(code int, msg string, data ...interface{}) (int, Reply) {
	if len(data) > 0 {
		return 200, Reply{
			Code: code,
			Msg:  msg,
			Data: data[0],
		}
	}
	return 200, Reply{
		Code: code,
		Msg:  msg,
	}
}

func Page(msg string, items interface{}, count int) (int, Reply) {
	return 200, Reply{
		Code: stSucc,
		Msg:  msg,
		Data: page{
			Items: items,
			Count: count,
		},
	}
}

// sucess
func Succ(msg string, data ...interface{}) (int, Reply) {
	return newReply(stSucc, msg, data...)
}

// jwt 错误
func ErrJwt(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrJwt, msg, data...)
}

//输入有误
func ErrIpt(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrIpt, msg, data...)
}

//输出错误
func ErrOpt(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrOpt, msg, data...)
}

//失败
func Fail(msg string, data ...interface{}) (int, Reply) {
	return newReply(stFail, msg, data...)
}
