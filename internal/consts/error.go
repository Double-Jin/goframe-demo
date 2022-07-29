//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package consts

import "fmt"

type ErrCode int

const (
	CodeOK                ErrCode = 200
	UnAuthorized          ErrCode = 401
	CodeNotFound          ErrCode = 404
	CodeInternalError     ErrCode = 500
	CodeUnknown           ErrCode = 520
	CodeValidationFailed  ErrCode = 9999
	CodeBusinessFailed    ErrCode = 10000
	AccountPassWordFailed ErrCode = 10001
	AccountLock           ErrCode = 10002
	CaptchaInvalid        ErrCode = 10003
	SqlError              ErrCode = 10004
	FileError              ErrCode = 10005
)

var errCodeName = map[ErrCode]string{
	CodeOK:                "Success",
	UnAuthorized:          "Not Authorized",
	CodeInternalError:     "INTERNAL ERROR",
	CodeNotFound:          "Resource Does Not Exist",
	CodeUnknown:           "Unknown Error",
	CodeValidationFailed:  "Validation Failed",
	AccountPassWordFailed: "账号或密码错误",
	AccountLock:           "账号已冻结",
	CodeBusinessFailed:    "业务出错",
	CaptchaInvalid:        "验证码输入错误",
	SqlError:              "查询错误",
	FileError:              "文件上传失败",
}

func (e ErrCode) Desc() string {
	if s, ok := errCodeName[e]; ok {
		return s
	}
	return fmt.Sprintf("unknown error code 0x%x", uint32(e))
}

func (e ErrCode) Code() int {
	return int(e)
}
