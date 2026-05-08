package repeat_rules

import "time"

type YearRule struct {
	interval DateInterval
}

func NewYearRule() YearRule {
	return YearRule{
		interval: DateInterval{Years: 1},
	}
}

func (r YearRule) CalcNextDate(now time.Time, start time.Time) time.Time {
	return CalcNextDateByInterval(now, start, r.interval)
}
