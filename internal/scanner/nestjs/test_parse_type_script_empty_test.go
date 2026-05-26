//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseTypeScript_Empty 테스트
package nestjs

import "testing"

func TestParseTypeScript_Empty(t *testing.T) {
	root, err := parseTypeScript([]byte(""))
	if err != nil {
		t.Fatal(err)
	}
	if root == nil {
		t.Fatal("expected non-nil root even for empty input")
	}
}
