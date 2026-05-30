//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestHandleFile_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestHandleFile_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := callExprFrom(t, `c.FormFile("upload")`)
	handleFile(ep, call)
	if ep.Request == nil || len(ep.Request.Files) != 1 || ep.Request.Files[0].Name != "upload" {
		t.Fatalf("file: %+v", ep.Request)
	}
}
