//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_JSON 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleResponse_JSON(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.Ident{Name: "result"},
		},
	}

	handleResponse(&ep, call, "json", nil, "handler")

	if len(ep.Responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(ep.Responses))
	}
	if ep.Responses[0].Status != "200" {
		t.Fatalf("expected status 200, got %s", ep.Responses[0].Status)
	}
	if ep.Responses[0].Kind != "json" {
		t.Fatalf("expected kind json, got %s", ep.Responses[0].Kind)
	}
}

func TestHandleResponse_SimpleKinds(t *testing.T) {
	for _, kind := range []string{"string", "data", "file"} {
		ep := scanner.Endpoint{}
		handleResponse(&ep, &ast.CallExpr{}, kind, nil, "handler")
		if len(ep.Responses) != 1 || ep.Responses[0].Status != "200" || ep.Responses[0].Kind != kind {
			t.Fatalf("%s: %v", kind, ep.Responses)
		}
	}
}

func TestHandleResponse_RedirectDefault(t *testing.T) {
	ep := scanner.Endpoint{}
	handleResponse(&ep, &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/x"`}}}, "redirect", nil, "handler")
	if ep.Responses[0].Status != "302" {
		t.Fatalf("redirect default status = %s", ep.Responses[0].Status)
	}
}

func TestHandleResponse_RedirectWithStatus(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{
		&ast.BasicLit{Kind: token.STRING, Value: `"/x"`},
		&ast.BasicLit{Kind: token.INT, Value: "301"},
	}}
	handleResponse(&ep, call, "redirect", nil, "handler")
	if ep.Responses[0].Status != "301" {
		t.Fatalf("redirect status = %s", ep.Responses[0].Status)
	}
}

func TestHandleResponse_StatusChainingSkipped(t *testing.T) {
	ep := scanner.Endpoint{}
	// c.Status(200) alone -> not recorded as standalone response
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "c"}, Sel: &ast.Ident{Name: "Status"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "200"}},
	}
	handleResponse(&ep, call, "status", nil, "handler")
	if len(ep.Responses) != 0 {
		t.Fatalf("c.Status() should not record a response, got %v", ep.Responses)
	}
}

func TestHandleResponse_SourceAndUnknownStatus(t *testing.T) {
	ep := scanner.Endpoint{}
	// status kind, SendStatus with no args -> status "(unknown)"; source set
	call := &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "c"}, Sel: &ast.Ident{Name: "SendStatus"}}}
	handleResponse(&ep, call, "status", nil, "respond")
	if ep.Responses[0].Status != "(unknown)" || ep.Responses[0].Source != "respond" {
		t.Fatalf("source/unknown: %+v", ep.Responses[0])
	}
}
