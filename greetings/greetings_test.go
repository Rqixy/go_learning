package greetings

import (
	"regexp"
	"testing"
)

// TestHelloNameはgreetings.Helloを名前で呼び出し、チェックします。
// 有効な戻り値の場合
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q nil`, msg, err, want)
	}
}

// TestHelloEmptyはgreetings.Helloをからの文字列で呼び出し、エラーをチェックします。
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
