package myerror

import "errors"

//根据业务逻辑需求,自定义一些错误
var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在。。。")
	ERROR_USER_EXITSTS   = errors.New("用户已经存在。。。")
	ERROR_USER_PWD       = errors.New("密码不正确")
)
