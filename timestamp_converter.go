package epoch

import (
	"time"
)

// ConvertSecondsToUTC converts a Unix timestamp in seconds to UTC datetime.
func ConvertSecondsToUTC(seconds int64) string {
	t := time.Unix(seconds, 0).UTC()
	return t.Format(time.RFC3339)
}

// ConvertMicrosecondsToUTC converts a Unix timestamp in microseconds to UTC datetime.
func ConvertMicrosecondsToUTC(microseconds int64) string {
	seconds := microseconds / 1e6
	nanoseconds := (microseconds % 1e6) * 1e3
	t := time.Unix(seconds, nanoseconds).UTC()
	return t.Format("2006-01-02T15:04:05.000000Z")
}
