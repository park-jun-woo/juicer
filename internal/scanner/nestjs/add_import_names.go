//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 단일 import 문에서 타입명과 경로를 결과 맵에 추가한다
package nestjs

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// addImportNames extracts type names from a single import statement and adds them to the result map.
func addImportNames(stmt *sitter.Node, src []byte, result map[string]string) {
	source := findChildByType(stmt, "string")
	if source == nil {
		return
	}
	path := unquoteTS(nodeText(source, src))
	if !strings.HasPrefix(path, ".") {
		return
	}
	clause := findChildByType(stmt, "import_clause")
	if clause == nil {
		return
	}
	named := findAllByType(clause, "import_specifier")
	for _, spec := range named {
		nameNode := findChildByType(spec, "identifier")
		if nameNode != nil {
			result[nodeText(nameNode, src)] = path
		}
	}
}
