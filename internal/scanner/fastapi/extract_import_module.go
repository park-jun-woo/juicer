//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what import_from_statement 에서 모듈 경로를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractImportModule extracts the module path from an import_from_statement.
// It checks relative_import first because `from .routes import router` has
// both a relative_import child (".routes") and a dotted_name child ("router"),
// and the module is the relative_import node, not the imported name.
func extractImportModule(stmt *sitter.Node, src []byte) string {
	relImport := findChildByType(stmt, "relative_import")
	if relImport != nil {
		return nodeText(relImport, src)
	}
	dotted := findChildByType(stmt, "dotted_name")
	if dotted != nil {
		return nodeText(dotted, src)
	}
	return ""
}
