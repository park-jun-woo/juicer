package scanner

import (
	"go/ast"
	"testing"
)

func TestResolveResponseType_NilInfoCase(t *testing.T) {
	tn, fields, conf := resolveResponseType(&ast.Ident{Name: "x"}, nil)
	if tn != "" || fields != nil || conf != "" {
		t.Fatal("expected empty")
	}
}
