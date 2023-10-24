package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello,World!")
	for _, s := range os.Environ() {
		fmt.Println(s)
	}
}
