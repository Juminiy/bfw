package lang

import "time"

func GetNS(duration time.Duration) float64 {
	return float64(duration / time.Nanosecond)
}

func GetUS(duration time.Duration) float64 {
	return float64(duration / time.Microsecond)
}

func GetMS(duration time.Duration) float64 {
	return float64(duration / time.Millisecond)
}

func GetS(duration time.Duration) float64 {
	return float64(duration / time.Second)
}

func GetMIN(duration time.Duration) float64 {
	return float64(duration / time.Minute)
}
