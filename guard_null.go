package guard

import (
	"fmt"

	"github.com/bgsrb/guard/messages"
)

func (arg Argument) NotNull() Argument {

	return arg.NotNullWithMeesage(messages.NotEmpty)
}

func (arg Argument) NotNullWithMeesage(message string) Argument {

	if arg.Error != nil {
		return arg
	}

	if arg.Value == nil {
		arg.Error = fmt.Errorf(message, arg.Name)
	}

	return arg
}
