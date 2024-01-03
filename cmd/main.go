package main

func addInt(a, b int) int {
	return a + b
}

func main() {
	addInt(1<<10, 1<<5)
	return
}
