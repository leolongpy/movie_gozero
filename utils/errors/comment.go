package errors

import (
	"errors"
	"fmt"
)

var (
	ErrorCommentFailed = errors.New(
		fmt.Sprintln("comment", "操作异常"),
	)
)
