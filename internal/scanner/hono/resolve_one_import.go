//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what 단일 import 구문에서 변수명→파일 경로 매핑을 수집한다
package hono

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func resolveOneImport(stmt *sitter.Node, src []byte, dir string, imports map[string]string, absRoot string) {
	pathNode := findChildByType(stmt, "string")
	if pathNode == nil {
		return
	}
	importPath := unquoteTS(nodeText(pathNode, src))
	if !strings.HasPrefix(importPath, ".") {
		return
	}
	resolved := resolveRelativePath(dir, importPath)
	if resolved == "" {
		return
	}
	clause := findChildByType(stmt, "import_clause")
	if clause == nil {
		return
	}
	ident := findChildByType(clause, "identifier")
	if ident != nil {
		imports[nodeText(ident, src)] = resolved
	}
	named := findChildByType(clause, "named_imports")
	if named != nil {
		collectNamedImports(named, src, resolved, imports)
	}
}
