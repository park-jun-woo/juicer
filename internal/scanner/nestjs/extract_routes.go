//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 컨트롤러 클래스에서 HTTP 라우트를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractControllers finds all @Controller-decorated classes and their routes.
func extractControllers(root *sitter.Node, src []byte, file string) []controllerInfo {
	var result []controllerInfo
	imports := extractImports(root, src)
	classes := findAllByType(root, "class_declaration")
	for _, cls := range classes {
		prefix, ok := controllerPrefix(cls, src)
		if !ok {
			continue
		}
		ci := controllerInfo{
			prefix:  prefix,
			imports: imports,
		}
		ci.endpoints = extractMethods(cls, src, file)
		result = append(result, ci)
	}
	return result
}
