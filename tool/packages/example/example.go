package example

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) Say() string {
	return fmt.Sprintf("name is %s age is %d", s.Name, s.Age)
}

type People interface {
	Say() string
}

func NewStudent(name string, age int) People {
	return &Student{
		Name: name,
		Age:  age,
	}
}
