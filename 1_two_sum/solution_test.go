package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		Nums   []int
		Target int
	}

	tests := []struct {
		Name string
		Args args
		Want []int
	}{
		{
			Name: "twoSum",
			Args: args{
				Nums:   []int{2, 7, 11, 15},
				Target: 9,
			},
			Want: []int{0, 1},
		},
		{
			Name: "twoSum",
			Args: args{
				Nums:   []int{3, 2, 4},
				Target: 6,
			},
			Want: []int{1, 2},
		},
		{
			Name: "twoSum",
			Args: args{
				Nums:   []int{3, 3},
				Target: 6,
			},
			Want: []int{0, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Want, twoSum(test.Args.Nums, test.Args.Target))
		})
	}

}
