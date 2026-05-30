//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCallFuncName_Attribute 테스트
package django

import "testing"

func TestCallFuncName_Attribute(t *testing.T) {
	src := []byte("x = views.health()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := firstCall(root)
	if call == nil {
		t.Fatal("no call")
	}
	if got := callFuncName(call, src); got != "views.health" {
		t.Fatalf("got %q, want views.health", got)
	}
}
