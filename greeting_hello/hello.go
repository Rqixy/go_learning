package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// あいさつメッセージを取得して出力します。
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}