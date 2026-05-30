//ff:func feature=scan type=test control=sequence
//ff:what scanBody — 핸들러 body 디스패치 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func handlerBody(t *testing.T, src string) *ast.BlockStmt {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Body != nil {
			return fn.Body
		}
	}
	t.Fatal("no func body")
	return nil
}

func TestScanBody_NilBody(t *testing.T) {
	ep := &scanner.Endpoint{}
	scanBody(ep, nil, "c", nil, nil, "handler")
	if ep.Request != nil {
		t.Fatal("nil body should be a no-op")
	}
}

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

func TestScanBody_ChainedStatusJSON(t *testing.T) {
	src := `package m
func h() {
	return c.Status(201).JSON(result)
}
`
	body := handlerBody(t, src)
	ep := &scanner.Endpoint{}
	scanBody(ep, body, "c", nil, &funcIndex{}, "handler")
	if len(ep.Responses) != 1 || ep.Responses[0].Status != "201" {
		t.Fatalf("chained status: %v", ep.Responses)
	}
}

func TestScanBody_ChainedBind(t *testing.T) {
	src := `package m
func h() {
	c.Bind().Body(req)
}
`
	body := handlerBody(t, src)
	ep := &scanner.Endpoint{}
	scanBody(ep, body, "c", nil, &funcIndex{}, "handler")
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatalf("chained Bind().Body should populate body: %+v", ep.Request)
	}
}

func TestScanBody_NonCtxCall(t *testing.T) {
	// a plain function call (not on ctx) with source "handler" -> 1-depth check
	src := `package m
func h() {
	doWork()
	other.Method()
}
`
	body := handlerBody(t, src)
	ep := &scanner.Endpoint{}
	// nil info -> checkOneDepthCall returns early; no panic
	scanBody(ep, body, "c", nil, &funcIndex{}, "handler")
	if ep.Request != nil || len(ep.Responses) != 0 {
		t.Fatalf("non-ctx calls should not populate ep: %+v", ep)
	}
}
