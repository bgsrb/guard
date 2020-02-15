package guard

type Argument struct {
	Err   error
	Name  string
	Value interface{}
}

func Guard(value interface{}, name string) Argument {
	return Argument{
		Name:  name,
		Value: value,
	}
}
