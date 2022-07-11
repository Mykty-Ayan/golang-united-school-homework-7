package coverage

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
const lenOneErrorString = "expected length is 1, but received %d"

func TestPeople_Len(t *testing.T) {
	p1 := Person{
		firstName: "Ayan",
		lastName:  "Akkassov",
		birthDay:  time.Date(2000, time.Month(2), 1, 4, 30, 30, 0, time.UTC),
	}
	people := make(People, 0, 3)

	people = append(people, p1)

	if len(people) != 1 {
		fmt.Println(people)
		t.Errorf(lenOneErrorString, len(people))
	}
	p2 := Person{
		"Madi",
		"Yussupov",
		time.Date(1997, time.Month(11), 1, 12, 32, 13, 0, time.UTC),
	}

	people = append(people, p2)

	if len(people) != 2 {
		fmt.Println(people)
		t.Errorf("enexpected lenght is 2, but received %d", len(people))
	}

	people = people[:len(people)-1]

	if len(people) != 1 {
		fmt.Println(people)
		t.Errorf(lenOneErrorString, len(people))
	}
}
