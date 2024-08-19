package helloDeps11

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps11.PrintPhrase")
	return phrase
}
