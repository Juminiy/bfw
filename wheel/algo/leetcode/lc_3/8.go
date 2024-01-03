package lc_3

func myAtoi(s string) int {
	const (
		i32Min int = 1 << 31
		i32Max int = (1 << 31) - 1
		i32Inf int = 1<<31 + 1
	)
	ans, sign, begin, end := 0, 0, false, false
	for _, ch := range s {
		switch ch {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			ans = opt(ans, ch)
			if ans > i32Inf {
				end = true
			}
			begin = true
		case '+', '-':
			if sign != 0 || begin {
				end = true
			} else {
				if ch == '-' {
					sign = -1
				} else {
					sign = 1
				}
				begin = true
			}
		case ' ':
			if begin {
				end = true
			}
		default:
			end = true
		}
		if end {
			break
		}
	}
	if sign == 0 {
		sign = 1
	}
	if sign == 1 && ans > i32Max {
		ans = i32Max
	} else if sign == -1 {
		if ans > i32Min {
			ans = -i32Min
		} else {
			ans = -ans
		}
	}
	return ans
}

func opt(i32 int, ch int32) int {
	return mul10(i32) + int(ch-'0')
}

func mul10(i32 int) int {
	return i32<<1 + i32<<3
}
