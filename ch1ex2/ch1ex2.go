package ch1ex2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Echo1() string {
	var s, sep string
	sep = "\n"

	for i := 0; i < len(os.Args); i++ {
		iS := strconv.Itoa(i)
		s += "index: " + iS + ", value: " + os.Args[i] + sep
	}
	fmt.Println(s)
	return s
}

func Echo2() string {
	s, sep := "", "\n"
	for i, arg := range os.Args[0:] {
		iS := strconv.Itoa(i)
		s += "index: " + iS + ", value: " + arg + sep
	}
	fmt.Println(s)
	return s
}

func Echo3() string {
	for i, x := range os.Args[0:] {
		iS := strconv.Itoa(i)
		os.Args[i] = "index: " + iS + ", value: " + x
	}

	s := fmt.Sprint(strings.Join(os.Args[0:], "\n"))
	fmt.Println(s)
	return s
}
