//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestExtractPathString_Round5 테스트
package echo

import "testing"

func TestExtractPathString_Round5(t *testing.T) {
	got, ok := extractPathString(nil, parseExpr(t, `"/users"`))
	if !ok || got != "/users" {
		t.Fatalf("literal: %q %v", got, ok)
	}

	got2, ok2 := extractPathString(nil, parseExpr(t, `"/a" + "/b"`))
	if !ok2 || got2 != "/a/b" {
		t.Fatalf("concat: %q %v", got2, ok2)
	}

	if _, ok := extractPathString(nil, parseExpr(t, `someVar`)); ok {
		t.Fatal("ident should not resolve without info")
	}
}
