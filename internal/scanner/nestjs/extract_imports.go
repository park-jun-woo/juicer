//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what import 문에서 타입명과 상대 경로 매핑을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractImports builds a map from imported type name to relative source path.
func extractImports(root *sitter.Node, src []byte) map[string]string {
	result := make(map[string]string)
	stmts := findAllByType(root, "import_statement")
	for _, stmt := range stmts {
		addImportNames(stmt, src, result)
	}
	return result
}
