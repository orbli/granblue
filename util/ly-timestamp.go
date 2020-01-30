package util

import (
	"time"
)

func Date(t time.Time) string {
	return t.Format("060102")
}

func Time(t time.Time) string {
	return t.Format("150405")
}

func Datetime(t time.Time) string {
	return t.Format("060102150405")
}
