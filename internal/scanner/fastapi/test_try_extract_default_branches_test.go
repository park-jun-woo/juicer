//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestTryExtractDefault_Branches 테스트
package fastapi

import "testing"

func TestTryExtractDefault_Branches(t *testing.T) {

	src := []byte("def f(x: int = 5): pass\n")
	root, _ := parsePython(src)

	if c := firstOfType(root, ":"); c != nil {
		v, call, none := tryExtractDefault(c, src)
		if v != "" || call != "" || none {
			t.Errorf(": node should yield empty, got %q %q %v", v, call, none)
		}
	}

	if c := firstOfType(root, "identifier"); c != nil {
		v, call, none := tryExtractDefault(c, src)
		if v != "" || call != "" || none {
			t.Errorf("identifier should yield empty, got %q %q %v", v, call, none)
		}
	}

	if c := firstOfType(root, "integer"); c != nil {
		v, _, none := tryExtractDefault(c, src)
		if v != "5" || none {
			t.Errorf("integer default: got %q none=%v", v, none)
		}
	}
}
