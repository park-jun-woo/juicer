//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what parseMethodBody 테스트 헬퍼
package sqls

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseMethodBody(t *testing.T, src, methodName string) *ast.BlockStmt {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Name.Name == methodName {
			return fn.Body
		}
	}
	t.Fatalf("method %s not found", methodName)
	return nil
}
