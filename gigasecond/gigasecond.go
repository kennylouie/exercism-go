package gigasecond

import (
	"time"
	"math"
)

func AddGigasecond(t time.Time) time.Time {
	gigasecond := math.Pow10(9)
	t = t.Add(time.Duration(gigasecond)*time.Second)
	return t
}
