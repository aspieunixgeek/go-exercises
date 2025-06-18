package main

import (
	"fmt"
	"github.com/aspieunixgeek/go-exercises/lenconv"
	"github.com/aspieunixgeek/go-exercises/tempconv"
	"github.com/aspieunixgeek/go-exercises/timeconv"
	"github.com/aspieunixgeek/go-exercises/wtconv"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		printLine()

		fah := tempconv.Fahrenheit(t)
		cel := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", fah, tempconv.FToC(fah), cel, tempconv.CToF(cel))

		ft := lenconv.Foot(t)
		mt := lenconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n", ft, lenconv.FtToMet(ft), mt, lenconv.MetToFt(mt))

		kg := wtconv.Kg(t)
		lb := wtconv.Lb(t)
		fmt.Printf("%s = %s, %s = %.24s\n", kg, wtconv.KgToLb(kg), lb, wtconv.LbToKg(lb))

		ps := timeconv.PicoSec(t)
		ns := timeconv.NanoSec(t)
		fmt.Printf("%s = %s, %s = %s\n", ps, timeconv.PsToNs(ps), ns, timeconv.NsToPs(ns))

		if i == len(os.Args[1:])-1 {
			printLine()
		}
	}
}

func printLine() {
	for i := range 32 {
		fmt.Print("--")
		if i == 31 {
			fmt.Print("\n")
		}
	}
}
