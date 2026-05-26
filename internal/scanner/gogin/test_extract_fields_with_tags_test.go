//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestExtractFields_WithTags 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestExtractFields_WithTags(t *testing.T) {
	src := `package test

type User struct {
	Name     string ` + "`json:\"name\" validate:\"required\"`" + `
	Email    string ` + "`json:\"email,omitempty\" binding:\"required\"`" + `
	Internal string ` + "`json:\"-\"`" + `
	NoTag    string
}

var u User
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}

	conf := types.Config{}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	_, err = conf.Check("test", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}

	var uIdent *ast.Ident
	for ident := range info.Defs {
		if ident.Name == "u" {
			uIdent = ident
			break
		}
	}

	_, fields := resolveExprType(uIdent, info)
	// Internal should be excluded (json:"-")
	// Expected: Name, Email, NoTag = 3 fields
	if len(fields) != 3 {
		t.Errorf("expected 3 fields (Internal excluded), got %d", len(fields))
		for _, f := range fields {
			t.Logf("  field: %+v", f)
		}
	}
}
