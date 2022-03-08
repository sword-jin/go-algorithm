package bitmap

func Bit1Count(num int64) int {
	count := 0
	for num != 0 {
		count += 1
		num = (num - 1) & num
	}
	return count
}

// Bit1Count2 通过不停的左移来进行比较
func Bit1Count2(num int64) int {
	count := 0
	var flag int64 = 1
	for flag != 0 {
		if num&flag != 0 {
			count += 1
		}
		flag <<= 1
	}
	return count
}

// Bit1Count1 在负数进来的情况下会死循环
func Bit1Count1(num int64) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count += 1
		}
		num >>= 1
	}
	return count
}
