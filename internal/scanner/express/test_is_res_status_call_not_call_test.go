//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResStatusCall_NotCall 테스트
package express

import "testing"

func TestIsResStatusCall_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`res;`))
	ids := findAllByType(fi.Root, "identifier")
	if isResStatusCall(ids[0], fi.Src) {
		t.Fatal("expected false")
	}
}
