package main

func add(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func getSub(a, b int) int {
	return a - b
}
