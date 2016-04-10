package main

import (
	"fmt"
	"io"
	"os"
)

func count(x int, m map[int]bool) {
	//	fmt.Println("X is ", x)
	for x > 0 {
		m[x%10] = true
		x /= 10
	}
	// fmt.Println("m is ", m)
}

func solve(t int, x int, out io.Writer) {
	if x == 0 {
		fmt.Fprintf(out, "Case #%d: INSOMNIA\n", t)
		return
	}
	m := make(map[int]bool)
	for i := 1; i <= 10000; i++ {
		count(x*i, m)
		if len(m) == 10 {
			fmt.Fprintf(out, "Case #%d: %d\n", t, x*i)
			return
		} else {
			//	fmt.Printf("Case %d", len(m))
		}
	}
	panic(x)
}

func solveAll() {
	out, _ := os.Create("qr1/a-complete.out")
	for i := 1; i < 100000; i++ {
		solve(i, i, out)
	}
}

func main() {
	in, err := os.Open("qr1/A-large.in")
	//in, err := os.Open("qr1/a-small.in")
	if err != nil {
		panic(err)
	}
	out, err := os.Create("qr1/A-large.out")

	var T int
	fmt.Fscanf(in, "%d", &T)

	fmt.Printf("T %d", T)
	for i := 0; i < T; i++ {
		var N int
		fmt.Fscanf(in, "%d", &N)
		solve(i+1, N, out)
	}
}
