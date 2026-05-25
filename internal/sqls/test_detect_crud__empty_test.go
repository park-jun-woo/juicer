//ff:func feature=sql type=parse control=sequence
//ff:what TestDetectCRUD_Empty 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestDetectCRUD_Empty(t *testing.T) {
	body := &ast.BlockStmt{}
	if detectCRUD(body) != "" {
		t.Fatal("expected empty")
	}
}
