package ch1ex1

import (
	"fmt"
	"os"
	"strings"
)

func Echo1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Echo3() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
