//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddWriteMethodBody_EmptyClass 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddWriteMethodBody_EmptyClass(t *testing.T) {
	ep := &scanner.Endpoint{}
	addWriteMethodBody(ep, "POST", "", nil)
	if ep.Request != nil {
		t.Fatal("expected no body for empty class")
	}
}
