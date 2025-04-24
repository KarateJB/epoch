package epoch

import (
	"time"
)

// ConvertEpochSecToUtc converts a Unix timestamp in seconds to UTC datetime.
func ConvertEpochSecToUtc(seconds int64) string {
	t := time.Unix(seconds, 0).UTC()
	return t.Format(time.RFC3339)
}

// ConvertEpochMsecToUtc converts a Unix timestamp in microseconds to UTC datetime.
func ConvertEpochMsecToUtc(microseconds int64) string {
	seconds := microseconds / 1e6
	nanoseconds := (microseconds % 1e6) * 1e3
	t := time.Unix(seconds, nanoseconds).UTC()
	return t.Format("2006-01-02T15:04:05.000000Z")
}

// ConvertUtcToEpochSec converts a UTC datetime string to epoch time in seconds.
func ConvertUtcToEpochSec(utc string) (int64, error) {
	t, err := time.Parse(time.RFC3339, utc)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// ConvertUtcToEpochMsec converts a UTC datetime string to epoch time in milliseconds.
func ConvertUtcToEpochMsec(utc string) (int64, error) {
	t, err := time.Parse(time.RFC3339, utc)
	if err != nil {
		return 0, err
	}
	return t.UnixNano() / 1e6, nil
}