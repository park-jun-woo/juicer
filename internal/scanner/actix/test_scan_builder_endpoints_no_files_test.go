//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestScanBuilderEndpoints_NoFiles 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestScanBuilderEndpoints_NoFiles(t *testing.T) {
	eps := scanBuilderEndpoints(nil, structIndex{}, map[string][]scanner.Field{}, map[string]*handlerInfo{})
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %+v", eps)
	}
}
