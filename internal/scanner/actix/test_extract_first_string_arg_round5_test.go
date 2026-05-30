//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractFirstStringArg_Round5 테스트
package actix

import "testing"

func TestExtractFirstStringArg_Round5(t *testing.T) {
	fi := aFi(t, `fn f() { web::resource("/users"); }`)
	call := aFirst(t, fi.root, "call_expression")
	if got := extractFirstStringArg(call, fi.src); got != "/users" {
		t.Fatalf("got %q", got)
	}
}
