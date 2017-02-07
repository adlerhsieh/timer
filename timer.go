package timer

import (
	"time"
)

func Now() string {
	t := time.Date(2017, 3, 1, 0, 0, 0, 0, time.UTC)
	if time.Now().Before(t) {
		return time.Now().UTC().Format("2006-01-02 15:04:05")
	} else {
		var err error
		panic(err)
	}
}
