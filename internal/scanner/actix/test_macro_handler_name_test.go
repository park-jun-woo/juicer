//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestMacroHandlerName 테스트
package actix

import "testing"

func TestMacroHandlerName(t *testing.T) {
	src := []byte(`async fn get_user() {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fn := findFuncByName(root, src, "get_user")
	if fn == nil {
		t.Fatal("function not found")
	}
	if got := macroHandlerName(fn, src); got != "get_user" {
		t.Fatalf("macroHandlerName = %q, want get_user", got)
	}
}
