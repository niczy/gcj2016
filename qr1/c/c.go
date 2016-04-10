package main

import (
	"fmt"
	"os"
	"strconv"
)

func module(s string, base int, m int) int {
	if len(s) == 1 {
		if s[len(s)-1] == '0' {
			return 0
		} else {
			return 1
		}
	}
	if s[len(s)-1] == '0' {
		return module(s[0:len(s)-1], base, m) * base % m
	} else {
		return (module(s[0:len(s)-1], base, m)*base + 1) % m
	}
}

func main() {
	//J := 50
	N := uint(16)
	//jj := 1
	out, _ := os.Create("qr1/c/c.out")
	fmt.Fprintf(out, "Case #1:\n")
	for i := 0; i <= ((1 << N) - 2); i++ {
		b := "1" + strconv.FormatInt(int64(i), 2) + "1"
		factors := make([]int, 0)
		for j := 2; j <= 10; j++ {
			for k := 2; k <= 100; k++ {
				if module(b, j, k) == 0 {
					factors = append(factors, k)
					break
				}
			}
		}
		if len(factors) == 9 {
			fmt.Println(b, factors)
		}
	}
}
