package epoch

import (
	"testing"
)

func TestConvertEpochSecToUtc(t *testing.T) {
	seconds := int64(1633072803) // Example timestamp in seconds
	expected := "2021-10-01T07:20:03Z"
	result := ConvertEpochSecToUtc(seconds)
	if result != expected {
		t.Errorf("ConvertEpochSecToUtc(%d) = %s; want %s", seconds, result, expected)
	}
}

func TestConvertEpochMsecToUtc(t *testing.T) {
	microseconds := int64(1633072800123456) // Example timestamp in microseconds
	expected := "2021-10-01T07:20:00.123456Z"
	result := ConvertEpochMsecToUtc(microseconds)
	if result != expected {
		t.Errorf("ConvertEpochMsecToUtc(%d) = %s; want %s", microseconds, result, expected)
	}
}

func TestConvertUtcToEpochMsec(t *testing.T) {
	utc := "2021-10-01T07:20:00.123Z" // Example UTC datetime string
	expected := int64(1633072800123)  // Expected epoch time in milliseconds
	result, err := ConvertUtcToEpochMsec(utc)
	if err != nil {
		t.Errorf("ConvertUtcToEpochMsec(%s) returned error: %v", utc, err)
	}
	if result != expected {
		t.Errorf("ConvertUtcToEpochMsec(%s) = %d; want %d", utc, result, expected)
	}
}

func TestConvertUtcToEpochSec(t *testing.T) {
	utc := "2021-10-01T07:20:03Z" // Example UTC datetime string
	expected := int64(1633072803) // Expected epoch time in seconds
	result, err := ConvertUtcToEpochSec(utc)
	if err != nil {
		t.Errorf("ConvertUtcToEpochSec(%s) returned error: %v", utc, err)
	}
	if result != expected {
		t.Errorf("ConvertUtcToEpochSec(%s) = %d; want %d", utc, result, expected)
	}
}
