package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("expected 1 argument")
		os.Exit(1)
	}
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("%d has %d non-zero bits\n", n, PopCount(n))
	fmt.Printf("%d = %b\n", n, n)
}

func PopCount(x uint64) int {
	result := uint64(0)
	for i := 0; i < 64; i++ {
		result += x & 1
		x >>= 1
	}
	return int(result)
}
