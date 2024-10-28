package utils

import "time"

// timeNowUTC3 + возвращает время +3
func TimeNowUTC3() time.Time {
	return time.Now().In(time.FixedZone("UTC+3", 3*60*60))
}
