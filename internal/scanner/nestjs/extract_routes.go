//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 컨트롤러 클래스에서 HTTP 라우트를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractControllers finds all @Controller-decorated classes and their routes.
// absFile is the absolute path of the source file (used to resolve cross-file
// imports for factory-pattern base controllers).
func extractControllers(root *sitter.Node, src []byte, file string, absFile string) []controllerInfo {
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
			version: controllerVersion(cls, src),
			imports: imports,
		}
		direct := extractMethods(cls, src, file)
		inherited := resolveBaseController(cls, src, absFile, imports, file)
		ci.endpoints = mergeEndpoints(inherited, direct)
		result = append(result, ci)
	}
	return result
}
