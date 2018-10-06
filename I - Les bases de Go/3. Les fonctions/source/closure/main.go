package main

func mafonction(a, b int) int {
	return a + b
}

func main() {
	c := 2

	maclosure := func(a, b int) int {
		return (a + b) * c
	}

	func(a, b int) int {
		return (a + b) * c
	}(9, 1)

	mafonction(1, 4)
	maclosure(1, 8)
}
