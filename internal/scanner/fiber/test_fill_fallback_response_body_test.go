//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestFillFallbackResponseBody — AST struct 인덱스로 응답 본문 폴백 해석 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestFillFallbackResponseBody(t *testing.T) {
	src := `package m
type Book struct {
	Title  string ` + "`json:\"title\"`" + `
	Rating int    ` + "`json:\"rating\"`" + `
	hidden string
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	idx := &funcIndex{astStructs: map[string]*ast.StructType{}}
	for _, decl := range file.Decls {
		indexStructDecl(decl, idx)
	}

	resp := &scanner.Response{Kind: "json", Body: "Book{}"}
	fillFallbackResponseBody(resp, idx)
	if resp.TypeName != "Book" {
		t.Fatalf("typeName: want Book, got %s", resp.TypeName)
	}
	if resp.Confidence != "partial" {
		t.Errorf("confidence: want partial, got %s", resp.Confidence)
	}
	if len(resp.Fields) != 2 {
		t.Fatalf("expected 2 exported fields, got %d", len(resp.Fields))
	}
	if resp.Fields[0].JSON != "title" || resp.Fields[1].JSON != "rating" {
		t.Errorf("json tags: got %+v", resp.Fields)
	}
}
