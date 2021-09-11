package errors

import (
	"errors"
	"fmt"
)

var (
	ErrorCinemaFailed = errors.New(
		fmt.Sprintln("cinema_rpc", "操作异常"),
	)
	ErrorCinemaNotFound = errors.New(
		fmt.Sprintln("cinema_rpc", "找不到对应的影院"),
	)
)
