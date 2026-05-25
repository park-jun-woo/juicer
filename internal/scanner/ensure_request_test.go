//ff:func feature=scan type=extract control=sequence
//ff:what TestEnsureRequest_NilRequest 테스트
package scanner

import "testing"

func TestEnsureRequest_NilRequest(t *testing.T) {
	ep := &Endpoint{}
	ensureRequest(ep)
	if ep.Request == nil {
		t.Fatal("expected non-nil request")
	}
}
