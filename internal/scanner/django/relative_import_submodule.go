//ff:func feature=scan type=extract control=sequence topic=django
//ff:what `from .X import ...` 문에서 단일 상대 서브모듈 이름을 추출한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// relativeImportSubmodule returns the single-level relative submodule name from
// a `from .X import ...` statement, or "" if it is not a single relative import.
func relativeImportSubmodule(stmt *sitter.Node, src []byte) string {
	rel := findChildByType(stmt, "relative_import")
	if rel == nil {
		return ""
	}
	dotted := findChildByType(rel, "dotted_name")
	if dotted == nil {
		return ""
	}
	name := nodeText(dotted, src)
	if name == "" || strings.Contains(name, ".") {
		return ""
	}
	return name
}
