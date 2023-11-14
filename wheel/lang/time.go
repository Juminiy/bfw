package lang

import "time"

func GetNS(duration time.Duration) int64 {
	return int64(duration / time.Nanosecond)
}

func GetUS(duration time.Duration) int64 {
	return int64(duration / time.Microsecond)
}

func GetMS(duration time.Duration) int64 {
	return int64(duration / time.Millisecond)
}

func GetS(duration time.Duration) int64 {
	return int64(duration / time.Second)
}

func GetMIN(duration time.Duration) int64 {
	return int64(duration / time.Minute)
}
