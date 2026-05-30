//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestAstFieldJSONName 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAstFieldJSONName(t *testing.T) {
	src := `package m
type S struct {
	Name string ` + "`json:\"name\"`" + `
	Skip string ` + "`json:\"-\"`" + `
	Plain string
}`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	var fields []*ast.Field
	ast.Inspect(file, func(n ast.Node) bool {
		if st, ok := n.(*ast.StructType); ok {
			fields = st.Fields.List
		}
		return true
	})
	if len(fields) < 3 {
		t.Fatalf("expected 3 fields, got %d", len(fields))
	}
	if got := astFieldJSONName(fields[0]); got != "name" {
		t.Fatalf("name: %q", got)
	}
	if got := astFieldJSONName(fields[1]); got != "" {
		t.Fatalf("dash should be empty: %q", got)
	}
	if got := astFieldJSONName(fields[2]); got != "" {
		t.Fatalf("no tag should be empty: %q", got)
	}
}
