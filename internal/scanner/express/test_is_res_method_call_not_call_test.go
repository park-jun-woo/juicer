//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResMethodCall_NotCall 테스트
package express

import "testing"

func TestIsResMethodCall_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`res;`))
	ids := findAllByType(fi.Root, "identifier")
	if _, ok := isResMethodCall(ids[0], fi.Src); ok {
		t.Fatal("expected false for non-call")
	}
}
