//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResMethodCall_StatusChain 테스트
package express

import "testing"

func TestIsResMethodCall_StatusChain(t *testing.T) {
	fi := mustParse(t, []byte(`res.status(201).json({});`))
	m, ok := isResMethodCall(outermostCall(fi), fi.Src)
	if !ok || m != "json" {
		t.Fatalf("got %q,%v", m, ok)
	}
}
