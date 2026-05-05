package repeat_rules

import (
	"fmt"
	"time"
)

type DayRule struct {
	interval DateInterval
}

const MaxDayRepeat = 400

func NewDayRule(daysCount int) (DayRule, error) {
	if daysCount > MaxDayRepeat {
		return DayRule{}, fmt.Errorf("days repeat interval should be less %d, got: %d", MaxDayRepeat, daysCount)
	}
	return DayRule{
		interval: DateInterval{Days: daysCount},
	}, nil

}

func (r DayRule) CalcNextDate(now time.Time, start time.Time) time.Time {
	return CalcNextDateByInterval(now, start, r.interval)
}
