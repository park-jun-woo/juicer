//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveAllModels 테스트
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestResolveAllModels(t *testing.T) {
	fields := []pydanticField{{name: "name", typeName: "str"}}
	files := []fileInfo{
		{absPath: "/a.py", models: map[string][]pydanticField{"User": fields}},
	}
	endpoints := []scanner.Endpoint{
		{Request: &scanner.Request{Body: &scanner.Body{TypeName: "User"}}},
	}
	reqs := []modelRequest{
		{typeName: "User", epIdx: 0, isBody: true},
	}
	resolveAllModels(reqs, endpoints, files)
	if len(endpoints[0].Request.Body.Fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(endpoints[0].Request.Body.Fields))
	}
}

func TestResolveAllModelsSkipBranches(t *testing.T) {
	fields := []pydanticField{{name: "name", typeName: "str"}}
	files := []fileInfo{
		{absPath: "/a.py", models: map[string][]pydanticField{"User": fields}},
	}
	endpoints := []scanner.Endpoint{
		{Request: &scanner.Request{Body: &scanner.Body{TypeName: "User"}}},
	}
	reqs := []modelRequest{
		// fields == nil: unknown model resolves to nil
		{typeName: "Unknown", epIdx: 0, isBody: true},
		// epIdx out of range: skipped
		{typeName: "User", epIdx: 99, isBody: true},
	}
	resolveAllModels(reqs, endpoints, files)
	if len(endpoints[0].Request.Body.Fields) != 0 {
		t.Fatalf("expected no fields applied, got %d", len(endpoints[0].Request.Body.Fields))
	}
}
