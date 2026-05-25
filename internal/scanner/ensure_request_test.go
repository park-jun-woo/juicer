package scanner

import "testing"

func TestEnsureRequest_NilRequest(t *testing.T) {
	ep := &Endpoint{}
	ensureRequest(ep)
	if ep.Request == nil {
		t.Fatal("expected non-nil request")
	}
}

func TestEnsureRequest_ExistingRequest(t *testing.T) {
	req := &Request{}
	ep := &Endpoint{Request: req}
	ensureRequest(ep)
	if ep.Request != req {
		t.Fatal("should not replace existing request")
	}
}
