//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCallFuncName_Neither 테스트
package django

import "testing"

func TestCallFuncName_Neither(t *testing.T) {

	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	if got := callFuncName(root, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
