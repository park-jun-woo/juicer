//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 단일 파일 AST에서 Annotated[T, Depends(func)] 형태의 별칭을 수집한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// collectFileAliases scans a single file AST for top-level assignments of the form
// X = Annotated[T, Depends(func)]. Returns alias -> depends func name.
func collectFileAliases(root *sitter.Node, src []byte) map[string]string {
	result := make(map[string]string)
	assignments := findAllByType(root, "assignment")
	for _, assign := range assignments {
		alias, fn := parseAnnotatedDependsAlias(assign, src)
		if alias != "" {
			result[alias] = fn
		}
	}
	return result
}
