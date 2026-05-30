//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestTryExtractDefault_Call 테스트
package fastapi

import "testing"

func TestTryExtractDefault_Call(t *testing.T) {
	src := []byte("def f(x = Query(default=5)): pass\n")
	root, _ := parsePython(src)
	c := firstOfType(root, "call")
	if c == nil {
		t.Skip("no call node")
	}
	v, call, none := tryExtractDefault(c, src)
	if call != "Query" || v == "" || none {
		t.Fatalf("call: v=%q call=%q none=%v", v, call, none)
	}
}
