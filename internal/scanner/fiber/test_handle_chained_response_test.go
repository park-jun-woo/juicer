//ff:func feature=scan type=test control=sequence
//ff:what TestHandleChainedResponse_StatusJSON 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleChainedResponse_StatusJSON(t *testing.T) {
	ep := scanner.Endpoint{}

	// c.Status(201)
	statusCall := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "c"},
			Sel: &ast.Ident{Name: "Status"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.INT, Value: "201"},
		},
	}

	// c.Status(201).JSON(result)
	outerCall := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   statusCall,
			Sel: &ast.Ident{Name: "JSON"},
		},
		Args: []ast.Expr{
			&ast.Ident{Name: "result"},
		},
	}

	handleChainedResponse(&ep, statusCall, outerCall, "JSON", nil, "handler")

	if len(ep.Responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(ep.Responses))
	}
	if ep.Responses[0].Status != "201" {
		t.Fatalf("expected status 201, got %s", ep.Responses[0].Status)
	}
	if ep.Responses[0].Kind != "json" {
		t.Fatalf("expected kind json, got %s", ep.Responses[0].Kind)
	}
	if ep.Responses[0].Body != "result" {
		t.Fatalf("expected body 'result', got %s", ep.Responses[0].Body)
	}
}

func chainCall(method string, args []ast.Expr) (*ast.CallExpr, *ast.CallExpr) {
	statusCall := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "c"}, Sel: &ast.Ident{Name: "Status"}},
		Args: args,
	}
	outerCall := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: statusCall, Sel: &ast.Ident{Name: method}},
		Args: nil,
	}
	return statusCall, outerCall
}

func TestHandleChainedResponse_Kinds(t *testing.T) {
	statusArg := []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "200"}}

	for _, tc := range []struct{ method, kind string }{
		{"SendString", "string"},
		{"Send", "data"},
		{"XML", "XML"}, // default branch
	} {
		ep := scanner.Endpoint{}
		sc, oc := chainCall(tc.method, statusArg)
		handleChainedResponse(&ep, sc, oc, tc.method, nil, "handler")
		if len(ep.Responses) != 1 || ep.Responses[0].Kind != tc.kind {
			t.Fatalf("%s: kind = %v", tc.method, ep.Responses)
		}
	}
}

func TestHandleChainedResponse_SourceAndUnknownStatus(t *testing.T) {
	ep := scanner.Endpoint{}
	// no status args -> status "(unknown)"; source != handler -> Source set
	sc, oc := chainCall("JSON", nil)
	handleChainedResponse(&ep, sc, oc, "JSON", nil, "respond")
	if len(ep.Responses) != 1 {
		t.Fatalf("expected 1 response")
	}
	if ep.Responses[0].Status != "(unknown)" {
		t.Errorf("status = %q, want (unknown)", ep.Responses[0].Status)
	}
	if ep.Responses[0].Source != "respond" {
		t.Errorf("source = %q, want respond", ep.Responses[0].Source)
	}
}
