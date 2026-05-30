//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddPathParams_Empty 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddPathParams_Empty(t *testing.T) {
	ep := &scanner.Endpoint{}
	addPathParams(ep, nil)
	if ep.Request != nil {
		t.Fatal("expected Request to stay nil for empty params")
	}
}
