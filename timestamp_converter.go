package epoch

import (
	"time"
)

// ConverSecondsToUtc converts a Unix timestamp in seconds to UTC datetime.
func ConverSecondsToUtc(seconds int64) string {
	t := time.Unix(seconds, 0).UTC()
	return t.Format(time.RFC3339)
}

// ConvertMicrosecondsToUtc converts a Unix timestamp in microseconds to UTC datetime.
func ConvertMicrosecondsToUtc(microseconds int64) string {
	seconds := microseconds / 1e6
	nanoseconds := (microseconds % 1e6) * 1e3
	t := time.Unix(seconds, nanoseconds).UTC()
	return t.Format("2006-01-02T15:04:05.000000Z")
}

// ConvertUtcToEpoch converts a UTC datetime string to epoch time in either milliseconds or seconds.
func ConvertUtcToEpoch(utc string, inMilliseconds bool) (int64, error) {
	t, err := time.Parse(time.RFC3339, utc)
	if err != nil {
		return 0, err
	}
	if inMilliseconds {
		return t.UnixNano() / 1e6, nil
	}
	return t.Unix(), nil
}
