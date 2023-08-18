package trendanalysis

import (
	"strconv"

	"github.com/go-gota/gota/series"
	"github.com/pkg/errors"
)

func Sum(s series.Series) (float64, error) {
	if s.Err != nil {
		return 0, errors.Wrap(s.Err, "error in Sum")
	}
	sum := 0.0
	for _, v := range s.Records() {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, errors.Wrapf(err, "unable to convert value '%s' to float", v)
		}
		sum += val
	}
	return sum, nil
}
