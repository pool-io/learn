package basics_test

import (
	"math/rand"
	"sort"
	"testing"

	basics "github.com/pool-io/learn/3-golang/0-basics"

	"github.com/stretchr/testify/assert"
)

const DEFAULT_LENGTH = 100

func TestReverse(t *testing.T) {
	input := make([]int, DEFAULT_LENGTH)
	expected := make([]int, DEFAULT_LENGTH)
	for i := 0; i < DEFAULT_LENGTH; i++ {
		num := rand.Intn(DEFAULT_LENGTH)

		input[i] = num
		expected[DEFAULT_LENGTH-1-i] = num
	}

	actual := basics.Reverse(input)

	for i, num := range actual {
		assert.Equal(t, expected[i], num)
	}
}

func TestSort(t *testing.T) {
	input := make([]int, DEFAULT_LENGTH)
	expected := make([]int, DEFAULT_LENGTH)
	for i := 0; i < DEFAULT_LENGTH; i++ {
		num := rand.Intn(DEFAULT_LENGTH)

		input[i] = num
		expected[i] = num
	}

	actual := basics.Sort(input)
	sort.Ints(expected)

	for i, num := range actual {
		assert.Equal(t, expected[i], num)
	}
}

var WORD_BANK = []string{
	"hello",
	"world",
	"foo",
	"bar",
	"pool",
	"stream",
	"tank",
	"drain",
	"caleb",
	"brian",
	"david",
	"isaac",
	"sam",
	"thomas",
	"soung bae",
	"cypress",
	"ycombinator",
	"hagwon",
	"uc berkeley",
	"uc irvine",
	"uc san diego",
	"uc santa barbara",
}

func TestSortNames(t *testing.T) {
	input := make([]string, DEFAULT_LENGTH)
	expected := make([]string, DEFAULT_LENGTH)
	for i := 0; i < DEFAULT_LENGTH; i++ {
		word := WORD_BANK[rand.Intn(len(WORD_BANK))]

		input[i] = word
		expected[i] = word
	}

	actual := basics.SortNames(input)
	sort.Strings(expected)

	for i, num := range actual {
		assert.Equal(t, expected[i], num)
	}
}

func TestNumCount(t *testing.T) {
	input := make([]int, DEFAULT_LENGTH)
	expected := make(map[int]int)

	for i := range input {
		num := rand.Intn(DEFAULT_LENGTH / 4)

		input[i] = num
		expected[num]++
	}

	actual := basics.NumCount(input)
	assert.Equal(t, expected, actual)
}
