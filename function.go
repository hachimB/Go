package main

func myFunc(callBack func(num int) int) int {
	return callBack(10)
}

func f(num int) int {
	return num * 2
}

func concatStr(str string) func(str2 string) string {
	return func(str2 string) string {
		return str + str2
	}
}

// variadic function

func sum(nums ...int) int {
	s := 0

	for _, val := range nums {
		s += val
	}
	return s
}

func mul(nums ...int) (m1 int, m2 int) {
	m1 = 1
	m2 = 1
	for _, m := range nums {
		m1 *= m
		m2 -= m
	}
	return m1, m2

}
