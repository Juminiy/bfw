package cal

import "fmt"

func run01(n int) {
	for i := 0; i*i <= n; i++ {
		for j := i; j*j <= n; j++ {
			if i*i+j*j == n {
				fmt.Println("x=", i, "y=", j)
			}
		}
	}
}
