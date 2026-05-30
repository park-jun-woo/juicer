//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCallFuncName_Identifier 테스트
package django

import "testing"

func TestCallFuncName_Identifier(t *testing.T) {
	src := []byte("x = path('a/', view)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := firstCall(root)
	if call == nil {
		t.Fatal("no call")
	}
	if got := callFuncName(call, src); got != "path" {
		t.Fatalf("got %q, want path", got)
	}
}
