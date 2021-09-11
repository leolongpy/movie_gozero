package errors

import (
	"errors"
	"fmt"
)

var (
	ErrorOrderFailed = errors.New(
		fmt.Sprintln("order", "操作异常"),
	)
	ErrorOrderAlreadyScore = errors.New(
		fmt.Sprintln("order", "已经评分了"),
	)
)
