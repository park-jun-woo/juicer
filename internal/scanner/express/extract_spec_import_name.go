//ff:func feature=scan type=extract control=sequence topic=express
//ff:what import_specifier에서 alias 또는 name을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractSpecImportName(spec *sitter.Node, src []byte) string {
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
