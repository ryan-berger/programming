package main

import "fmt"

func main() {
	var name string
	fmt.Scanln(&name)

	result := ""
	result += string(rune(name[0]))
	lastLetter := name[0]

	for i := 1; i < len(name); i++ {
		if name[i] == lastLetter {
			for j := i; j < len(name); j++ {
				if name[j] == lastLetter {
					i++
				} else {
					goto FINISH
				}
			}
		}

	FINISH:
		if i < len(name) {
			result += string(rune(name[i]))
			lastLetter = name[i]
		}
	}

	fmt.Println(result)
}
