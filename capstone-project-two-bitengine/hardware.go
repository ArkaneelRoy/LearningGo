package main

const maxInt64 int64 = 1<<63 - 1

func divMod(a, b int64) (int64, int64, bool) {
	if b == 0 {
		return 0, 0, false
	}
	return a / b, a % b, true
}

func safePow(base, exp int64) (int64, bool) {
	if exp < 0 {
		return 0, false
	}
	if exp == 0 {
		return 1, true
	}
	var res int64 = 1
	for i := int64(0); i < exp; i++ {
		if base > 0 && res > maxInt64/base {
			return 0, false
		}
		res *= base
	}
	return res, true
}
