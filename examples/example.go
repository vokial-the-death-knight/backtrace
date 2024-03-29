package main

import (
	"fmt"

	"github.com/vokial-the-death-knight/backtrace"
)

func main() {
	fmt.Println("main")
	f1()
}

func f1() {
	fmt.Println("f1")
	f2()
}

func f2() {
	fmt.Println("f2")
	fmt.Println(backtrace.DebugPrintBacktrace(backtrace.DebugBacktrace()))
}
