package calendar

import "time"

// JulianDate2000 the Julian day number for the day starting at 12:00 on January 1, 2000.
const JulianDate2000 = 2451545

// ZeroDate the zero Julian date converted to time.Time.
var ZeroDate = time.Date(-4713, 11, 24, 12, 0, 0, 0, time.UTC)

// TimeToJulianDate converts time.Time to Julian date.
func TimeToJulianDate(t time.Time) float64 {
	month := int(t.Month())
	a := (14 - month) / 12
	y := t.Year() + 4800 - a
	m := month + 12 * a - 3
	julianDays := t.Day() + (153 * m + 2) / 5 + 365 * y + y / 4 - y / 100 + y / 400 - 32045
	julianTime := (float64(t.Hour()) - 12) / 24 + float64(t.Minute()) / 1440 + float64(t.Second()) / 86400
	return float64(julianDays) + julianTime
}

// TimeToJulianDaysAndTime converts time.Time to Julian days (an integer value) and Julian time (a fractional number).
func TimeToJulianDaysAndTime(t time.Time) (float64, float64) {
	julianDate := DateTimeToJulianDate(t)
	julianDays := float64(int(julianDate))
	return julianDays, julianDate - julianDays
}