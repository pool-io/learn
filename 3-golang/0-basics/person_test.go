package basics_test

import (
	"testing"

	basics "github.com/pool-io/learn/3-golang/0-basics"
	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {
	name := "name"
	age := 10

	person := basics.NewPerson(name, age)

	assert.Equal(t, name, basics.Name(person))
	assert.Equal(t, age, basics.Age(person))

	newName := "new name"
	newAge := 12

	renamed := basics.SetName(person, newName)
	assert.Equal(t, newName, basics.Name(renamed))
	assert.Equal(t, age, basics.Age(renamed))

	aged := basics.SetAge(person, newAge)
	assert.Equal(t, name, basics.Name(aged))
	assert.Equal(t, newAge, basics.Age(aged))
}
