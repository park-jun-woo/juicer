//ff:func feature=sql type=test control=sequence
//ff:what TestParseMethodDecl_NoReceiver 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestParseMethodDecl_NoReceiver(t *testing.T) {
	fn := &ast.FuncDecl{Name: &ast.Ident{Name: "Foo"}, Type: &ast.FuncType{}}
	if parseMethodDecl(fn) != nil {
		t.Fatal("expected nil")
	}
}
