package common

import "time"

func NowUTC() time.Time {
	return time.Now().UTC()
}

func NowUTCUnix() int64 {
	return NowUTC().Unix()
}
