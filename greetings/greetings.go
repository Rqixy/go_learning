package greetings

import "fmt"

// Helloは、指定された人のあいさつを返します。
func Hello(name string) string {
	// メッセージに名前を埋め込む挨拶を返します。
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	var message string = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}