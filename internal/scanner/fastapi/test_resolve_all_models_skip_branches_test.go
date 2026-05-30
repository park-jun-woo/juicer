//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveAllModelsSkipBranches 테스트
package fastapi

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveAllModelsSkipBranches(t *testing.T) {
	fields := []pydanticField{{name: "name", typeName: "str"}}
	files := []fileInfo{
		{absPath: "/a.py", models: map[string][]pydanticField{"User": fields}},
	}
	endpoints := []scanner.Endpoint{
		{Request: &scanner.Request{Body: &scanner.Body{TypeName: "User"}}},
	}
	reqs := []modelRequest{

		{typeName: "Unknown", epIdx: 0, isBody: true},

		{typeName: "User", epIdx: 99, isBody: true},
	}
	resolveAllModels(reqs, endpoints, files)
	if len(endpoints[0].Request.Body.Fields) != 0 {
		t.Fatalf("expected no fields applied, got %d", len(endpoints[0].Request.Body.Fields))
	}
}
