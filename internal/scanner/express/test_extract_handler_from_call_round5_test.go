//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractHandlerFromCall_Round5 테스트
package express

import "testing"

func TestExtractHandlerFromCall_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`x = wrap(handler);`))
	call := exFirst(t, fi, "call_expression")
	got := extractHandlerFromCall(call, fi.Src)
	if got == "" {
		t.Fatalf("expected handler name, got %q", got)
	}
}
