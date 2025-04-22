# Epoch Package

The `epoch` package provides utility functions to convert timestamps into UTC date-time strings. It supports conversions for timestamps in seconds and microseconds.

## Functions

### `ConverSecondsToUtc(seconds int64) string`
Converts a timestamp in seconds since the Unix epoch to a UTC date-time string in ISO 8601 format.

**Example:**
```go
seconds := int64(1633072800)
result := ConverSecondsToUtc(seconds)
// result: "2021-10-01T00:00:00Z"
```

### `ConvertMicrosecondsToUtc(microseconds int64) string`
Converts a timestamp in microseconds since the Unix epoch to a UTC date-time string in ISO 8601 format with microsecond precision.

**Example:**
```go
microseconds := int64(1633072800123456)
result := ConvertMicrosecondsToUtc(microseconds)
// result: "2021-10-01T00:00:00.123456Z"
```

## Running Unit Tests

To run the unit tests for the `epoch` package, use the following command:

```bash
go test ./...
```
