package util

import "time"

func ConvertStringToTime(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s)
	return t, err
}

func ConvertTimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}
