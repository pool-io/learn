package structs_test

import (
	"testing"

	structs "github.com/pool-io/learn/3-golang/1-structs"
	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {
	name := "name"
	age := 10

	person := structs.NewPerson(name, age)

	assert.Equal(t, name, structs.Name(person))
	assert.Equal(t, age, structs.Age(person))

	newName := "new name"
	newAge := 12

	renamed := structs.SetName(person, newName)
	assert.Equal(t, newName, structs.Name(renamed))
	assert.Equal(t, age, structs.Age(renamed))

	aged := structs.SetAge(person, newAge)
	assert.Equal(t, name, structs.Name(aged))
	assert.Equal(t, newAge, structs.Age(aged))
}
