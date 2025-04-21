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

func TestConvertUTCToEpochMilliseconds(t *testing.T) {
	utc := "2021-10-01T07:20:00.123Z" // Example UTC datetime string
	expected := int64(1633072800123) // Expected epoch time in milliseconds
	result, err := ConvertUTCToEpoch(utc, true)
	if err != nil {
		t.Errorf("ConvertUTCToEpoch(%s, true) returned error: %v", utc, err)
	}
	if result != expected {
		t.Errorf("ConvertUTCToEpoch(%s, true) = %d; want %d", utc, result, expected)
	}
}

func TestConvertUTCToEpochSeconds(t *testing.T) {
	utc := "2021-10-01T07:20:03Z" // Example UTC datetime string
	expected := int64(1633072803) // Expected epoch time in seconds
	result, err := ConvertUTCToEpoch(utc, false)
	if err != nil {
		t.Errorf("ConvertUTCToEpoch(%s, false) returned error: %v", utc, err)
	}
	if result != expected {
		t.Errorf("ConvertUTCToEpoch(%s, false) = %d; want %d", utc, result, expected)
	}
}
