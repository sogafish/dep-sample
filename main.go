package main

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"runtime"
)

func main() {
	scheduler.Every(10).Seconds().Run(hello)

	runtime.Goexit()
}

func hello() {
	fmt.Println("Oh....")
}
