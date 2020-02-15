package guard

import (
	"errors"
	"fmt"

	"github.com/bgsrb/guard/messages"
)

func (arg Argument) MinLength(minLength int) Argument {

	if arg.Err != nil {
		return arg
	}

	switch arg.Value.(type) {
	case string:
		if len(arg.Value.(string)) < minLength {
			arg.Err = fmt.Errorf(messages.MinLength, arg.Name, minLength)
		}
		break
	default:
		arg.Err = errors.New("The value's type is invalid")
		break
	}

	return arg
}
