package decisionanalysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyingRecommendation(t *testing.T) {
	output := GetRecommendation(40000, 36, 0.03, 15000, nil, nil, nil, nil, nil)
	assert.Equal(t, "Buying is recommended.", output, "Buying is recommended.")
}
