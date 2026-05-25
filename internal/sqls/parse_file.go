//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what 단일 *_repo.go 파일을 AST 파싱하여 메서드 스켈레톤 추출
package sqls

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// parseFile parses a single *_repo.go file and extracts method skeletons.
func parseFile(path string) ([]MethodSkeleton, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}

	var results []MethodSkeleton

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if sk := parseMethodDecl(fn); sk != nil {
			results = append(results, *sk)
		}
	}

	return results, nil
}
