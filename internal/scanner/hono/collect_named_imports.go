//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what named_imports 노드에서 변수명→파일 경로 매핑을 수집한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func collectNamedImports(named *sitter.Node, src []byte, resolved string, imports map[string]string) {
	for _, spec := range childrenOfType(named, "import_specifier") {
		nameNode := spec.ChildByFieldName("name")
		if nameNode != nil {
			imports[nodeText(nameNode, src)] = resolved
		}
	}
}
