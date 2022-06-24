package greetings

import (
	"errors"
	"fmt"
)

// Helloは、指定された人のあいさつを返します。
func Hello(name string) (string, error) {
	// 名前が指定されていない場合は、メッセージと共にエラーを返します。
	if name == "" {
		return "", errors.New("empty name")
	}

	// メッセージに名前を埋め込む挨拶を返します。
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	var message string = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
