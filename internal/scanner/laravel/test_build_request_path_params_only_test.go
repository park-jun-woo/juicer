//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestBuildRequest_PathParamsOnly 테스트
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestBuildRequest_PathParamsOnly(t *testing.T) {
	pp := []scanner.Param{{Name: "id", Type: "integer"}}
	r := buildRequest(t.TempDir(), pp, nil, map[string]*fileInfo{})
	if r == nil || len(r.PathParams) != 1 || r.Body != nil {
		t.Fatalf("got %+v", r)
	}
}
