package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello は、指定された人のあいさつを返します。
func Hello(name string) (string, error) {
	// 名前が指定されていない場合は、メッセージと共にエラーを返します。
	if name == "" {
		return "", errors.New("empty name")
	}

	// メッセージに名前を埋め込む挨拶を返します。
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	//var message string = fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf(randomFormat())
	return message, nil
}

// Hellos は、指定された各人を関連つけるマップを返します。
// あいさつ付きメッセージ
func Hellos(names []string) (map[string]string, error) {
	// 名前をメッセージに関連つけるためのマップ
	messages := make(map[string]string)

	// 受信した名前のスライスをループし、呼び出します。
	// 各名前のメッセージを取得するHello関数
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	// マップで、取得したメッセージを名前に関連つけます
	return messages, nil
}

// Initは、関数で使用される変数の初期値を設定します。
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormatは、一連のgreetingMessageの1つを返します。
// 返されてメッセージはランダムに選択されます。
func randomFormat() string {
	// メッセージ形式のスライス
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you %v",
		"Hail, %v! Well met!",
	}

	// 指定してランダムに選択されたメッセージ形式を返します。
	// フォーマットのスライスのランダムインデックス
	return formats[rand.Intn(len(formats))]
}
