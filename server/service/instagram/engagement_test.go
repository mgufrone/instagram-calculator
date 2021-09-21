package instagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvgCalculator(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		likes       []int64
		comments    []int64
		followers   int64
		expectedAvg float64
	}{
		{
			[]int64{0, 0},
			[]int64{0, 0},
			100,
			0.00,
		},
		{
			[]int64{1, 2, 3},
			[]int64{1, 2, 3},
			100,
			float64(12) / float64(3 * 100),
		},
		{
			[]int64{5, 15, 10},
			[]int64{0, 1, 20},
			100,
			float64(51) / float64(3 * 100),
		},
	}
	for _, c := range testCases {
		eng := &Engagement{
			Likes:     c.likes,
			Comments:  c.comments,
			Followers: c.followers,
		}
		assert.Equal(t, c.expectedAvg, eng.Rate())
	}
}
