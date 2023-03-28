package stats

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateStdDev(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}

	stddev, err := CalculateStdDev(x)
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.NotNil(t, stddev, "Expected stddev to not be nil")
	assert.NotEqual(t, 0, int(stddev), "Expected stddev not be zero as int")
	assert.NotEqual(t, 0.0, stddev, "Expected stddev not be zero as float64")
}
