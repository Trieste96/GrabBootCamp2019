package prime

import (
	"fmt"
	"testing"
	"golang/mock"
	"github.com/stretchr/testify/assert"
)

func TestCheckPrimeNumber(t *testing.T) {
	testTable := []struct {
		number   int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
	}
	assert := assert.New(t)
	for _, testCase := range testTable {
		fmt.Printf("Number: %d\n", testCase.number)
		assert.Equal(testCase.expected, checkPrimeNumber(testCase.number), "Failed at number: %d", testCase.number)
	}
}
