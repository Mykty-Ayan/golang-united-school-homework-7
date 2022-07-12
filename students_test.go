package coverage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

	if people.Len() != 1 {
		fmt.Println(people)
		t.Errorf(lenOneErrorString, len(people))
	}
	p2 := Person{
		"Madi",
		"Yussupov",
		time.Date(1997, time.Month(11), 1, 12, 32, 13, 0, time.UTC),
	}

	people = append(people, p2)

	if people.Len() != 2 {
		fmt.Println(people)
		t.Errorf("enexpected lenght is 2, but received %d", len(people))
	}

	people = people[:people.Len()-1]

	if people.Len() != 1 {
		fmt.Println(people)
		t.Errorf(lenOneErrorString, people.Len())
	}
}

func TestPeople_Less(t *testing.T) {

	people := People{
		Person{
			firstName: "Ayan",
			lastName:  "Akkassov",
			birthDay:  time.Date(2000, time.Month(2), 1, 4, 30, 30, 0, time.UTC),
		},
		Person{
			firstName: "John",
			lastName:  "Larry",
			birthDay:  time.Date(2000, time.Month(2), 1, 4, 30, 30, 0, time.UTC),
		},
		Person{
			firstName: "John",
			lastName:  "Barry",
			birthDay:  time.Date(2000, time.Month(2), 1, 4, 30, 30, 0, time.UTC),
		},
		Person{
			firstName: "Jane",
			lastName:  "Barry",
			birthDay:  time.Date(2000, time.Month(2), 2, 4, 30, 30, 0, time.UTC),
		},
	}

	tData := []struct {
		i        int
		j        int
		Expected bool
	}{
		{
			i:        0,
			j:        1,
			Expected: true,
		},
		{
			i:        1,
			j:        2,
			Expected: false,
		},
		{
			i:        1,
			j:        1,
			Expected: false,
		},
		{
			i:        2,
			j:        3,
			Expected: false,
		},
	}

	for _, tc := range tData {
		got := people.Less(tc.i, tc.j)
		assert.Equal(t, tc.Expected, got, "Less with %s, %s expected %t, but got %t", people[tc.i], people[tc.j], tc.Expected, got)
	}

}
