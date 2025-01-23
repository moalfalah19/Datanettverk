package stringer

import "fmt"

type Student struct {
	ID    int
	Name  string
	Age   int
}

func (s Student) String() string {
	return fmt.Sprintf("Student ID: %d. Name: %s. Age: %d.", s.ID, s.Name, s.Age)
}

