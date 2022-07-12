package goutil

import (
	"testing"
	"time"
)

func TestTimeConversion(t *testing.T) {
	t.Run("TimeToUnix", func(t *testing.T) {
		time := time.Now()
		unix := TimeToUnix(time)
		if unix != time.Unix() {
			t.Errorf("TimeToUnix(%v) = %v, want %v", time, unix, time.Unix())
		}
	})

	t.Run("UnixToTime", func(t *testing.T) {
		time := time.Now()
		unix := TimeToUnix(time)
		if unix != time.Unix() {
			t.Errorf("TimeToUnix(%v) = %v, want %v", time, unix, time.Unix())
		}
	})

	t.Run("TimeToMillis", func(t *testing.T) {
		time := time.Now()
		millis := TimeToMillis(time)
		if millis != time.UnixNano()/1e6 {
			t.Errorf("TimeToMillis(%v) = %v, want %v", time, millis, time.UnixNano()/1e6)
		}
	})

	t.Run("MillisToTime", func(t *testing.T) {
		time := time.Now()
		millis := TimeToMillis(time)
		if millis != time.UnixNano()/1e6 {
			t.Errorf("TimeToMillis(%v) = %v, want %v", time, millis, time.UnixNano()/1e6)
		}
	})

	t.Run("TimeToMicros", func(t *testing.T) {
		time := time.Now()
		micros := TimeToMicros(time)
		if micros != time.UnixNano()/1e3 {
			t.Errorf("TimeToMicros(%v) = %v, want %v", time, micros, time.UnixNano()/1e3)
		}
	})

	t.Run("MicrosToTime", func(t *testing.T) {
		time := time.Now()
		micros := TimeToMicros(time)
		if micros != time.UnixNano()/1e3 {
			t.Errorf("TimeToMicros(%v) = %v, want %v", time, micros, time.UnixNano()/1e3)
		}
	})

	t.Run("TimeToNanos", func(t *testing.T) {
		time := time.Now()
		nanos := TimeToNanos(time)
		if nanos != time.UnixNano() {
			t.Errorf("TimeToNanos(%v) = %v, want %v", time, nanos, time.UnixNano())
		}
	})

	t.Run("NanosToTime", func(t *testing.T) {
		time := time.Now()
		nanos := TimeToNanos(time)
		if nanos != time.UnixNano() {
			t.Errorf("TimeToNanos(%v) = %v, want %v", time, nanos, time.UnixNano())
		}
	})
}

func TestTime(t *testing.T) {
	time := time.Now()

	t.Run("timezone", func(t *testing.T) {
		zone, offset := GetTimezone(time)
		if zone != "UTC" {
			t.Errorf("timezone: %s, want UTC", zone)
		}

		if offset != 0 {
			t.Errorf("offset: %d, want 0", offset)
		}
	})

	t.Run("Change timezone", func(t *testing.T) {
		if offset := ConvertTimezone(time, "Europe/London").Hour(); offset != time.Hour()+1 {
			t.Errorf("ConvertTimezone(%v, %v) = %v, want %v", time, "Europe/London", offset, time.Hour()+1)
		}
		if offset := ConvertTimezone(time, "Asia/Tokyo").Hour(); offset != time.Hour()+9 {
			t.Errorf("ConvertTimezone(%v, %v) = %v, want %v", time, "Asia/Tokyo", offset, time.Hour()+9)
		}
	})
}
