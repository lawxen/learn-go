package main

import "fmt"

func Reverse(s string) string {
    fmt.Printf("input: %q\n", s)
    r := []rune(s)
    fmt.Printf("runes: %q\n", r)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}


func main() {
	input := "This quick brown fox jumped overlazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n",input)
	fmt.Printf("reversed: %q\n",rev)
	fmt.Printf("reversed again: %q\n",doubleRev)
}
