package main

import (
	"fmt"
	"io"
	"os"
)

func count(bytes []byte, goal byte) int {
	var ret int
	if len(bytes) == 1 {
		if bytes[0] == goal {
			return 0
		} else {
			return 1
		}
	}
	if bytes[len(bytes)-1] == goal {
		ret = count(bytes[0:len(bytes)-1], goal)
	} else {
		ret = 1 + count(bytes[0:len(bytes)-1], flipByte(goal))
	}
	return ret
}

func flipByte(b byte) byte {
	if b == '+' {
		return '-'
	} else {
		return '+'
	}
}

func flip(bytes []byte) []byte {
	for i := 0; i < (len(bytes))/2; i++ {
		tmp := flipByte(bytes[len(bytes)-i-1])
		bytes[len(bytes)-i-1] = flipByte(bytes[i])
		bytes[i] = tmp
	}
	if len(bytes)%2 == 1 {
		bytes[len(bytes)/2] = flipByte(bytes[len(bytes)/2])
	}

	return bytes
}

func solve(t int, out io.Writer, s string) {
	fmt.Fprintf(out, "Case #%d: %d\n", t, count([]byte(s), '+'))
}

func main() {
	in, err := os.Open("qr1/b/B-large.in")
	//in, err := os.Open("qr1/b/b.in")
	if err != nil {
		panic(err)
	}
	out, err := os.Create("qr1/b/B-large.out")
	var T int
	fmt.Fscanf(in, "%d", &T)
	for i := 1; i <= T; i++ {
		var s string
		fmt.Fscanf(in, "%s", &s)
		fmt.Println(s)
		solve(i, out, s)
	}
}
