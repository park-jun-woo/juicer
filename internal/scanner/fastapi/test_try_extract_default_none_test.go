//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestTryExtractDefault_None 테스트
package fastapi

import "testing"

func TestTryExtractDefault_None(t *testing.T) {
	src := []byte("def f(x: int = None): pass\n")
	root, _ := parsePython(src)
	c := firstOfType(root, "none")
	if c == nil {
		t.Skip("no none node found")
	}
	if _, _, none := tryExtractDefault(c, src); !none {
		t.Fatal("none node should set isNone")
	}
}
