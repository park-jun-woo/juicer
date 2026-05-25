package sqls

import (
	"go/ast"
	"testing"
)

func TestDetectCRUD_Nil(t *testing.T) {
	if detectCRUD(nil) != "" {
		t.Fatal("expected empty")
	}
}

func TestDetectCRUD_Empty(t *testing.T) {
	body := &ast.BlockStmt{}
	if detectCRUD(body) != "" {
		t.Fatal("expected empty")
	}
}
