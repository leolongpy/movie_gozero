package errors

import (
	"errors"
	"fmt"
)

const (
	errorCodeFilmSuccess = 200
)

var (
	ErrorFilmFailed = errors.New(
		fmt.Sprintln("film", "操作异常"),
	)
)
