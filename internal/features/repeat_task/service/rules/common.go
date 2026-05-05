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

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func CalcNextDateByInterval(now, start time.Time, interval DateInterval) time.Time {
	date := start
	for {
		date = date.AddDate(interval.Years, interval.Months, interval.Days)
		if afterNow(date, now) {
			break
		}
	}
	return date
}
