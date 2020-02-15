package guard

import (
	"fmt"

	"github.com/bgsrb/guard/messages"
)

func (arg Argument) NotEmpty() Argument {

	return arg.NotEmptyWithMeesage(messages.NotEmpty)
}

func (arg Argument) NotEmptyWithMeesage(message string) Argument {

	if arg.Error != nil {
		return arg
	}

	if arg.Value == "" {
		arg.Error = fmt.Errorf(message, arg.Name)
	}

	return arg
}
