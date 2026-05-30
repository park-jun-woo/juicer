//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestHandleForm_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestHandleForm_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := callExprFrom(t, `c.FormValue("title")`)
	handleForm(ep, call)
	if ep.Request == nil || len(ep.Request.FormFields) != 1 || ep.Request.FormFields[0].Name != "title" {
		t.Fatalf("form: %+v", ep.Request)
	}
	handleForm(ep, call)
	if len(ep.Request.FormFields) != 1 {
		t.Fatalf("dup: %+v", ep.Request.FormFields)
	}
}
