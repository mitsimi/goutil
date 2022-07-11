package goutil

import "time"

// Here is a summary of the components of a layout string.
// Each element shows by example the formatting of an element of the reference time.
// Only these values are recognized. Text in the layout string that is not recognized
// as part of the reference time is echoed verbatim during Format
// and expected to appear verbatim in the input to Parse.
//
// Year: "2006" "06"
// Month: "Jan" "January"
// Textual day of the week: "Mon" "Monday"
// Numeric day of the month: "2" "_2" "02"
// Numeric day of the year: "__2" "002"
// Hour: "15" "3" "03" (PM or AM)
// Minute: "4" "04"
// Second: "5" "05"
// AM/PM mark: "PM"

// Numeric time zone offsets format as follows:
// "-0700"  ±hhmm
// "-07:00" ±hh:mm
// "-07"    ±hh
//
// Replacing the sign in the format with a Z triggers the ISO 8601 behavior of printing Z instead of an offset for the UTC zone. Thus:
// "Z0700"  Z or ±hhmm
// "Z07:00" Z or ±hh:mm
// "Z07"    Z or ±hh

// Some valid layouts are invalid time values for time.Parse,
// due to formats such as _ for space padding and Z for zone information.
const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"

	// Log = "2006-01-02 15:04:05"

	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)

// TimeToUnix converts time.Time to Unix timestamp
func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}

// UnixToTime converts Unix timestamp to time.Time
func UnixToTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// TimeToMillis converts time.Time to milliseconds
func TimeToMillis(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

//	MillisToTime converts milliseconds to time.Time
func MillisToTime(millis int64) time.Time {
	return time.Unix(0, millis*1e6)
}

// TimeToMicros converts time.Time to microseconds
func TimeToMicros(t time.Time) int64 {
	return t.UnixNano() / 1e3
}

//	MicrosToTime converts microseconds to time.Time
func MicrosToTime(micros int64) time.Time {
	return time.Unix(0, micros*1e3)
}

// TimeToNanos converts time.Time to nanoseconds
func TimeToNanos(t time.Time) int64 {
	return t.UnixNano()
}

//	NanosToTime converts nanoseconds to time.Time
func NanosToTime(nanos int64) time.Time {
	return time.Unix(0, nanos)
}

// Timezone returns the timezone and offset of time.Time
// The offset is the difference between UTC and local time in seconds.
func GetTimezone(t time.Time) (string, int) {
	return t.Zone()
}

// ConvertTimezone converts time.Time to another timezone
func ConvertTimezone(t time.Time, zone string) time.Time {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return t
	}
	return t.In(loc)
}
