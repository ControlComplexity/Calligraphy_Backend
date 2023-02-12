package utils

import "time"

func FormatTimeToString(t time.Time) string{
	return t.Format("2006-01-02")
}
