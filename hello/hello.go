package hello

import (
	"fmt"
	"time"
)

func Hello() string {
	fmt.Println("hello world")
	return "Hello, world."
}

func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
