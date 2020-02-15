package guard

type Argument struct {
	Error error
	Name  string
	Value interface{}
}

func Guard(value interface{}, name string) Argument {
	return Argument{
		Name:  name,
		Value: value,
	}
}
