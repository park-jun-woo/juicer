//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what named import에서 변수명을 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractNamedImportVarName(clause *sitter.Node, src []byte) string {
	named := findChildByType(clause, "named_imports")
	if named == nil {
		return ""
	}
	spec := findChildByType(named, "import_specifier")
	if spec == nil {
		return ""
	}
	if alias := spec.ChildByFieldName("alias"); alias != nil {
		return nodeText(alias, src)
	}
	if name := spec.ChildByFieldName("name"); name != nil {
		return nodeText(name, src)
	}
	if first := findChildByType(spec, "identifier"); first != nil {
		return nodeText(first, src)
	}
	return ""
}
