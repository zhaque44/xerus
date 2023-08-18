package trendanalysis

import (
	"testing"

	"github.com/go-gota/gota/series"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	data := []string{"1.5", "2.5", "3.5", "4.5", "5.5"}
	s := series.New(data, series.Float, "test")

	result, err := Sum(s)
	if err != nil {
		t.Errorf("Unexpected error in Sum aggregation: %v", err)
	}

	expectedSum := 1.5 + 2.5 + 3.5 + 4.5 + 5.5

	assert.Equal(t, expectedSum, result)
}
