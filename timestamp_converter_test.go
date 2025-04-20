package epoch

import (
	"testing"
)

func TestConvertSecondsToUTC(t *testing.T) {
	seconds := int64(1633072803) // Example timestamp in seconds
	expected := "2021-10-01T07:20:03Z"
	result := ConvertSecondsToUTC(seconds)
	if result != expected {
		t.Errorf("ConvertSecondsToUTC(%d) = %s; want %s", seconds, result, expected)
	}
}

func TestConvertMicrosecondsToUTC(t *testing.T) {
	microseconds := int64(1633072800123456) // Example timestamp in microseconds
	expected := "2021-10-01T07:20:00.123456Z"
	result := ConvertMicrosecondsToUTC(microseconds)
	if result != expected {
		t.Errorf("ConvertMicrosecondsToUTC(%d) = %s; want %s", microseconds, result, expected)
	}
}
