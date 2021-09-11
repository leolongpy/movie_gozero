package errors

import (
	"errors"
	"fmt"
)

var (
	ErrorUserSuccess = errors.New(

		fmt.Sprintln("user", "操作成功"),
	)

	ErrorUserFailed = errors.New(

		fmt.Sprintln("user", "操作异常"),
	)

	ErrorUserAlready = errors.New(

		fmt.Sprintln("user", "该邮箱已经被注册过了"),
	)

	ErrorUserLoginFailed = errors.New(

		fmt.Sprintln("user", "密码或者用户名错误"),
	)

	ErrorScoreForbid = errors.New(
		fmt.Sprintln("user", "你没有买过该电影票，无法进行评分"),
	)
)
