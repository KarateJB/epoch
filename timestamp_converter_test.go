package epoch

import (
	"testing"
)

func TestConvertSecondsToUtc(t *testing.T) {
	seconds := int64(1633072803) // Example timestamp in seconds
	expected := "2021-10-01T07:20:03Z"
	result := ConverSecondsToUtc(seconds)
	if result != expected {
		t.Errorf("ConverSecondsToUtc(%d) = %s; want %s", seconds, result, expected)
	}
}

func TestConvertMicrosecondsToUtc(t *testing.T) {
	microseconds := int64(1633072800123456) // Example timestamp in microseconds
	expected := "2021-10-01T07:20:00.123456Z"
	result := ConvertMicrosecondsToUtc(microseconds)
	if result != expected {
		t.Errorf("ConvertMicrosecondsToUtc(%d) = %s; want %s", microseconds, result, expected)
	}
}

func TestConvertUtcToEpochMilliseconds(t *testing.T) {
	utc := "2021-10-01T07:20:00.123Z" // Example UTC datetime string
	expected := int64(1633072800123)  // Expected epoch time in milliseconds
	result, err := ConvertUtcToEpoch(utc, true)
	if err != nil {
		t.Errorf("ConvertUtcToEpoch(%s, true) returned error: %v", utc, err)
	}
	if result != expected {
		t.Errorf("ConvertUtcToEpoch(%s, true) = %d; want %d", utc, result, expected)
	}
}

func TestConvertUtcToEpochSeconds(t *testing.T) {
	utc := "2021-10-01T07:20:03Z" // Example UTC datetime string
	expected := int64(1633072803) // Expected epoch time in seconds
	result, err := ConvertUtcToEpoch(utc, false)
	if err != nil {
		t.Errorf("ConvertUtcToEpoch(%s, false) returned error: %v", utc, err)
	}
	if result != expected {
		t.Errorf("ConvertUtcToEpoch(%s, false) = %d; want %d", utc, result, expected)
	}
}
