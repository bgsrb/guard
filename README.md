# Guard
Guard is a fluent argument validation library.

# Usage
### Without Guard
```go
type Person struct {
 	Name string
	Age  int
}

func NewPerson(name string, age int) (Person, error) {

	if name == nil {
		return Person{}, errors.New("Name cannot be null.")
	}

	if name == "" {
		return Person{}, errors.New("Name cannot be empty.")
	}

	if len(name) < 3 {
		return Person{}, errors.New("Name cannot be shorter than 3 characters")
	}

	return Person{
		Name: name,
		Age:  age,
	}, nil

}
```
### With Guard
```go
type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) (Person, error) {

	v := Guard(name, "Name").NotNull().NotEmpty().MinLength(3)
	
	if v.Error != nil {
		return Person{}, v.Error
	} else {
		return Person{
			Name: name,
			Age:  age,
		}, nil
	}

}
```
