package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		a, b, sum int
	}{
		{
			a:   1,
			b:   1,
			sum: 2,
		},
		{
			a:   1,
			b:   -1,
			sum: 0,
		},
		{
			a:   -1,
			b:   -1,
			sum: -2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.sum, Sum(test.a, test.b))
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		a, abs int
	}{
		{
			a:   1,
			abs: 1,
		},
		{
			a:   -1,
			abs: 1,
		},
		{
			a:   0,
			abs: 0,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.abs, Abs(test.a))
	}
}

func TestProduct(t *testing.T) {
	tests := []struct {
		a, b, product int
	}{
		{
			a:       1,
			b:       1,
			product: 1,
		},
		{
			a:       1,
			b:       -1,
			product: -1,
		},
		{
			a:       -1,
			b:       -1,
			product: 1,
		},
		{
			a:       0,
			b:       1,
			product: 0,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.product, Product(test.a, test.b))
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		a, b, floor int
	}{
		{
			a:     1,
			b:     1,
			floor: 1,
		},
		{
			a:     1,
			b:     -1,
			floor: -1,
		},
		{
			a:     -1,
			b:     -1,
			floor: 1,
		},
		{
			a:     5,
			b:     2,
			floor: 2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.floor, Floor(test.a, test.b))
	}
}

func TestQuotient(t *testing.T) {
	tests := []struct {
		a, b     int
		quotient float64
	}{
		{
			a:        1,
			b:        1,
			quotient: 1,
		},
		{
			a:        1,
			b:        -1,
			quotient: -1,
		},
		{
			a:        -1,
			b:        -1,
			quotient: 1,
		},
		{
			a:        5,
			b:        2,
			quotient: 2.5,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.quotient, Quotient(test.a, test.b))
	}
}

func TestRemainder(t *testing.T) {
	tests := []struct {
		a, b, remainder int
	}{
		{
			a:         1,
			b:         1,
			remainder: 0,
		},
		{
			a:         1,
			b:         -1,
			remainder: 0,
		},
		{
			a:         -1,
			b:         -1,
			remainder: -2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.remainder, Remainder(test.a, test.b))
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		a, b, sum int
	}{
		{
			a:   1,
			b:   1,
			sum: 1,
		},
		{
			a:   2,
			b:   2,
			sum: 4,
		},
		{
			a:   2,
			b:   0,
			sum: 1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.sum, Power(test.a, test.b))
	}
}

func TestCube(t *testing.T) {
	tests := []struct {
		a, cube int
	}{
		{
			a:    1,
			cube: 1,
		},
		{
			a:    2,
			cube: 8,
		},
		{
			a:    10,
			cube: 1000,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.cube, Cube(test.a))
	}
}
