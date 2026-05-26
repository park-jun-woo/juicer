//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what applyModelFields 테스트
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestApplyModelFields(t *testing.T) {
	fields := []scanner.Field{{Name: "name", Type: "string"}}

	// body branch
	ep := &scanner.Endpoint{
		Request:   &scanner.Request{Body: &scanner.Body{}},
		Responses: []scanner.Response{{Kind: "json"}},
	}
	applyModelFields(ep, true, fields)
	if len(ep.Request.Body.Fields) != 1 {
		t.Fatalf("expected 1 body field, got %d", len(ep.Request.Body.Fields))
	}

	// response branch
	ep2 := &scanner.Endpoint{
		Responses: []scanner.Response{{Kind: "json"}},
	}
	applyModelFields(ep2, false, fields)
	if len(ep2.Responses[0].Fields) != 1 {
		t.Fatalf("expected 1 response field, got %d", len(ep2.Responses[0].Fields))
	}

	// no-op: body but no request
	ep3 := &scanner.Endpoint{}
	applyModelFields(ep3, true, fields)

	// no-op: response but no responses
	ep4 := &scanner.Endpoint{}
	applyModelFields(ep4, false, fields)
}
