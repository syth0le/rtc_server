package main

import (
	"fmt"
	"runtime"
)

func main() {
	version := runtime.Version()
	result := fmt.Sprintf("go version specification for the project: %s", version)
	fmt.Println(result)
	fmt.Scanf("h")
}
