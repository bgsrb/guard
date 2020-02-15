package guard

import (
	"fmt"

	"github.com/bgsrb/guard/messages"
)

func (arg Argument) NotNull() Argument {

	if arg.Err != nil {
		return arg
	}

	if arg.Value == nil {
		arg.Err = fmt.Errorf(messages.NotEmpty, arg.Name)
	}

	return arg
}
