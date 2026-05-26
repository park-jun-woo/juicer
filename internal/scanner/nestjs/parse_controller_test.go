//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseTypeScript_Valid 테스트
package nestjs

import "testing"

func TestParseTypeScript_Valid(t *testing.T) {
	src := []byte(`const x: number = 42;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	if root == nil {
		t.Fatal("expected non-nil root")
	}
}
