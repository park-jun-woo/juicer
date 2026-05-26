//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what import_from_statement 에서 import된 이름 목록을 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractImportNames extracts the imported names from an import_from_statement.
func extractImportNames(stmt *sitter.Node, src []byte) []string {
	var names []string
	for i := 0; i < int(stmt.ChildCount()); i++ {
		child := stmt.Child(i)
		name := tryImportedName(child, stmt, i, src)
		if name != "" {
			names = append(names, name)
		}
	}
	return names
}
