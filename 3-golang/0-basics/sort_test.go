package basics_test

import (
	"math/rand"
	"testing"

	basics "github.com/pool-io/learn/3-golang/0-basics"

	"github.com/stretchr/testify/assert"
)

const LENGTH_OF_NUMS = 100

func TestReverse(t *testing.T) {
	input := make([]int, LENGTH_OF_NUMS)
	expected := make([]int, LENGTH_OF_NUMS)
	for i := 0; i < LENGTH_OF_NUMS; i++ {
		num := rand.Intn(LENGTH_OF_NUMS)

		input[i] = num
		expected[LENGTH_OF_NUMS-1-i] = num
	}

	actual := basics.Reverse(input)

	for i, num := range actual {
		assert.Equal(t, expected[i], num)
	}
}

func TestSort(t *testing.T) {
	input := make([]int, LENGTH_OF_NUMS)
	expected := make([]int, LENGTH_OF_NUMS)
	for i := 0; i < LENGTH_OF_NUMS; i++ {
		num := rand.Intn(LENGTH_OF_NUMS)

		input[i] = num
		expected[LENGTH_OF_NUMS-1-i] = num
	}
}
