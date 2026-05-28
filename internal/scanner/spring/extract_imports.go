//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what Java import 문에서 타입명과 패키지 경로를 수집한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractImports(root *sitter.Node, src []byte) map[string]string {
	result := make(map[string]string)
	imports := findAllByType(root, "import_declaration")
	for _, imp := range imports {
		name, fqcn := parseOneImport(imp, src)
		if name != "" {
			result[name] = fqcn
		}
	}
	return result
}
