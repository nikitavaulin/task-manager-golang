package repeat_rules

import (
	"time"
)

type RepeatRule interface {
	CalcNextDate(now time.Time, start time.Time) time.Time
}

type DateInterval struct {
	Years  int
	Months int
	Days   int
}

func CalcNextDateByInterval(now, start time.Time, interval DateInterval) time.Time {
	date := start
	for date.Before(now) {
		date = date.AddDate(interval.Years, interval.Months, interval.Days)
	}
	return date
}
