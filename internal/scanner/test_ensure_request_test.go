//ff:func feature=scan type=extract control=sequence
//ff:what TestEnsureRequest 테스트
package scanner

import (
	"testing"
)

func TestEnsureRequest(t *testing.T) {
	ep := &Endpoint{}
	EnsureRequest(ep)
	if ep.Request == nil {
		t.Error("expected request to be created")
	}
	// Second call should not overwrite
	ep.Request.RawBody = true
	EnsureRequest(ep)
	if !ep.Request.RawBody {
		t.Error("expected existing request to be preserved")
	}
}
