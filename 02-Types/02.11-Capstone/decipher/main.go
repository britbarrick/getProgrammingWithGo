package main

import "fmt"

func main() {
	hiddenMessage := "CSOITEUIWUIZNSROCNKFD"
	key := "GOLANG"
	message := ""
	keyInd := 0

	for i := 0; i < len(hiddenMessage); i++ {
		m := hiddenMessage[i] - 'A'
		k := key[keyInd] - 'A'

		m = (m-k+26)%26 + 'A'
		message += string(m)

		keyInd++
		keyInd %= len(key)
	}
	fmt.Println(message)

}
