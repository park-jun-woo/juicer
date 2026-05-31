//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 컨트롤러 클래스에서 HTTP 라우트를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractControllers finds all @Controller-decorated classes and their routes.
// absFile is the absolute path of the source file (used to resolve cross-file
// imports for factory-pattern base controllers). projectRoot is the scan root,
// used to resolve enum member-expression paths across files.
func extractControllers(root *sitter.Node, src []byte, file string, absFile string, projectRoot string) []controllerInfo {
	var result []controllerInfo
	imports := extractImports(root, src)
	classes := findAllByType(root, "class_declaration")
	for _, cls := range classes {
		ci, ok := buildControllerInfo(cls, src, file, absFile, imports, root, projectRoot)
		if !ok {
			continue
		}
		result = append(result, ci)
	}
	return result
}
