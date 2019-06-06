package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(cipherMe("this is how it goes", "turtle"))
}

func cipherMe(text string, key string) string {
	txt := strings.ToUpper(strings.Replace(text, " ", "", -1))
	keyCap := strings.ToUpper(strings.Replace(key, " ", "", -1))
	keyInd := 0
	newMessage := ""

	for i := 0; i < len(txt); i++ {
		m := txt[i] + 'A'
		k := keyCap[keyInd] + 'A'

		m = (m+k)%26 + 'A'
		newMessage += string(m)

		keyInd++
		keyInd %= len(keyCap)
	}
	return newMessage
}
