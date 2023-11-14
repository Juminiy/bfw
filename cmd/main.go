package main

import (
	"bfw/wheel/num"
	"os"
	"strconv"
)

func main() {
	var (
		maxBit      int = 22
		eachBitLoop int = 1
	)
	if argsLen := len(os.Args); argsLen >= 2 {
		switch argsLen {
		case 2:
			{
				maxBit, _ = strconv.Atoi(os.Args[1])
			}
		case 3:
			{
				eachBitLoop, _ = strconv.Atoi(os.Args[2])
			}
		}
	}
	num.RunBigNumberMultiply(maxBit, eachBitLoop)
}
