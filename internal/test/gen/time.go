package gen

import "time"

func Now() time.Time {
	return time.Now().Truncate(time.Microsecond)
}
