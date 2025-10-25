// Use map to count the frequency of each character in a string and return the most frequent character.

// add a use case of counting the frequency of each character in a string by using a map
package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	mapS := make(map[byte]int)
	maxCount := 0
	var maxChar byte
	for i := 0; i < len(s); i++ {
		mapS[s[i]]++
		if mapS[s[i]] > maxCount {
			maxCount = mapS[s[i]]
			maxChar = s[i]
		}
	}
	fmt.Printf("%c\n",maxChar)
}