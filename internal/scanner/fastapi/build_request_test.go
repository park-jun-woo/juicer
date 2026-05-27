//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what buildRequest 테스트
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestBuildRequest(t *testing.T) {
	// With path params
	ri := routeInfo{
		params: []scanner.Param{{Name: "id", Type: "integer"}},
	}
	req := buildRequest(ri)
	if req == nil || len(req.PathParams) != 1 {
		t.Fatal("expected request with path params")
	}

	// With body
	ri2 := routeInfo{bodyType: "UserCreate"}
	req2 := buildRequest(ri2)
	if req2 == nil || req2.Body == nil {
		t.Fatal("expected request with body")
	}

	// With files
	ri3 := routeInfo{files: []scanner.Param{{Name: "f", Type: "file"}}}
	req3 := buildRequest(ri3)
	if req3 == nil || len(req3.Files) != 1 {
		t.Fatal("expected request with files")
	}

	// Empty
	ri4 := routeInfo{}
	req4 := buildRequest(ri4)
	if req4 != nil {
		t.Fatal("expected nil request for empty")
	}
}
