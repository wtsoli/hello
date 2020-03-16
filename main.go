// string & rune compare,
package main

import "fmt"
import "github.com/wtsoli/hello"

// string & rune compare,
func stringAndRuneCompare() {
	// string,
	s := "hello你好"

	fmt.Printf("%s, type: %T, len: %d\n", s, s, len(s))
	fmt.Printf("s[%d]: %v, type: %T\n", 0, s[0], s[0])
	li := len(s) - 1 // last index,
	fmt.Printf("s[%d]: %v, type: %T\n\n", li, s[li], s[li])

	// []rune
	rs := []rune(s)
	fmt.Printf("%v, type: %T, len: %d\n", rs, rs, len(rs))
	hello.Say(s)
}

func main() {
	stringAndRuneCompare()
}
