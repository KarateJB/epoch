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

func TestConvertUTCToMilliseconds(t *testing.T) {
	utc := "2021-10-01T07:20:00.123Z" // Example UTC datetime string
	expected := int64(1633072800123) // Expected epoch time in milliseconds
	result, err := ConvertUTCToMilliseconds(utc)
	if err != nil {
		t.Errorf("ConvertUTCToMilliseconds(%s) returned error: %v", utc, err)
	}
	if result != expected {
		t.Errorf("ConvertUTCToMilliseconds(%s) = %d; want %d", utc, result, expected)
	}
}

func TestConvertUTCToSeconds(t *testing.T) {
	utc := "2021-10-01T07:20:03Z" // Example UTC datetime string
	expected := int64(1633072803) // Expected epoch time in seconds
	result, err := ConvertUTCToSeconds(utc)
	if err != nil {
		t.Errorf("ConvertUTCToSeconds(%s) returned error: %v", utc, err)
	}
	if result != expected {
		t.Errorf("ConvertUTCToSeconds(%s) = %d; want %d", utc, result, expected)
	}
}
