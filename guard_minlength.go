package guard

import (
	"errors"
	"fmt"

	"github.com/bgsrb/guard/messages"
)

func (arg Argument) MinLength(minLength int) Argument {

	return arg.MinLengthWithMeesage(minLength, messages.MinLength)
}

func (arg Argument) MinLengthWithMeesage(minLength int, message string) Argument {

	if arg.Error != nil {
		return arg
	}

	switch arg.Value.(type) {
	case string:
		if len(arg.Value.(string)) < minLength {
			arg.Error = fmt.Errorf(message, arg.Name, minLength)
		}
		break
	default:
		arg.Error = errors.New("The value's type is invalid")
		break
	}

	return arg
}
