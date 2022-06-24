package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// ログエントリプレフィックスと時間、ソースファイル、行番号
	// 印刷を無効にするフラグなど、事前定義されたロガーのプロパティを設定します。
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	// あいさつメッセージを取得して出力します。
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	// エラーが返されない場合は、
	// 返されたメッセージをコンソールに出力します
	fmt.Println(message)
}
