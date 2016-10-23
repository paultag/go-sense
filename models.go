package sense

import (
	"strconv"
	"time"
)

type senseTime struct {
	time.Time
}

func (t *senseTime) UnmarshalJSON(buf []byte) error {
	i, err := strconv.ParseInt(string(buf), 10, 64)
	if err != nil {
		return err
	}
	tm := time.Unix(i/1000, 0)
	t.Time = tm
	return nil
}

type TimelineEvent struct {
	timezoneOffset int
	validActions   []string
	Timestamp      senseTime
	Type           string `json:"event_type"`
	Message        string
	SleepState     string
	durationMillis int
	SleepDepth     int
}

type Timeline struct {
	Events         []TimelineEvent
	Message        string
	ScoreCondition string `json:"score_condition"`
}
