//ff:func feature=scan type=extract control=sequence
//ff:what TestEnsureRequest_ExistingRequest 테스트
package scanner

import "testing"

func TestEnsureRequest_ExistingRequest(t *testing.T) {
	req := &Request{}
	ep := &Endpoint{Request: req}
	EnsureRequest(ep)
	if ep.Request != req {
		t.Fatal("should not replace existing request")
	}
}
