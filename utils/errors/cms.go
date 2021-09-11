package errors

import (
	"errors"
	"fmt"
)

var (
	ErrorCMSFailed = errors.New(
		fmt.Sprintln("cms", "操作异常"),
	)
	ErrorCMSLogin = errors.New(
		fmt.Sprintln("cms", "登录异常"),
	)
	ErrorCMSFailedParam = errors.New(
		fmt.Sprintln("cms", "参数异常"),
	)
	ErrorCMSForbiddenParam = errors.New(
		fmt.Sprintln("cms", "没有查询的权限"),
	)
	ErrorCMSAlreadyRegister = errors.New(
		fmt.Sprintln("cms", "已经添加过影院"),
	)
)
