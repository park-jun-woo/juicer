//ff:func feature=scan type=test control=sequence
//ff:what TestScanBody_RequestAndResponse 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestScanBody_RequestAndResponse(t *testing.T) {
	src := `package m
func h() {
	c.BodyParser(&req)
	_ = c.Query("page")
	_ = c.Params("id")
	_ = c.FormValue("title")
	_, _ = c.FormFile("avatar")
	_ = c.Body()
	return c.JSON(result)
}
`
	body := handlerBody(t, src)
	ep := &scanner.Endpoint{}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}
	scanBody(ep, body, "c", nil, idx, "handler")

	if ep.Request == nil {
		t.Fatal("expected request populated")
	}
	if ep.Request.Body == nil {
		t.Error("BodyParser not handled")
	}
	if len(ep.Request.Query) != 1 {
		t.Error("Query not handled")
	}
	if len(ep.Request.PathParams) != 1 {
		t.Error("Params not handled")
	}
	if len(ep.Request.FormFields) != 1 {
		t.Error("FormValue not handled")
	}
	if len(ep.Request.Files) != 1 {
		t.Error("FormFile not handled")
	}
	if !ep.Request.RawBody {
		t.Error("Body (raw) not handled")
	}
	if len(ep.Responses) != 1 || ep.Responses[0].Kind != "json" {
		t.Errorf("JSON response not handled: %v", ep.Responses)
	}
}
