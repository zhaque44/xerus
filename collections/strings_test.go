package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToIntSliceWithValidation(t *testing.T) {
	testCases := []struct {
		input    string
		expected []int
		errMsg   string
	}{
		{"1,2,3,4,5", []int{1, 2, 3, 4, 5}, ""},
		{"  42  ,  0, -7 ", []int{42, 0, -7}, ""},
		{"", nil, ""},
		{"abc,def,ghi", nil, "strconv.Atoi: parsing \"abc\": invalid syntax"},
		{"1, 2, 3, xyz, 4", nil, "strconv.Atoi: parsing \"xyz\": invalid syntax"},
	}

	for _, tc := range testCases {
		result, err := IntSliceFromString(tc.input)

		assert.Equal(t, tc.expected, result)

		if tc.errMsg == "" {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tc.errMsg)
		}
	}
}
