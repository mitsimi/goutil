package goutil

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	time := time.Now()

	t.Run("timezone", func(t *testing.T) {
		zone, offset := GetTimezone(time)
		if zone != "CEST" {
			t.Errorf("timezone: %s, want CEST", zone)
		}

		if offset != 7200 {
			t.Errorf("offset: %d, want 2", offset/3600)
		}
	})

	t.Run("Change timezone", func(t *testing.T) {
		if offset := ConvertTimezone(time, "Europe/London").Hour(); offset != time.Hour()-1 {
			t.Errorf("ConvertTimezone(%v, %v) = %v, want %v", time, "Europe/London", offset, time.Hour()-1)
		}
	})
}
