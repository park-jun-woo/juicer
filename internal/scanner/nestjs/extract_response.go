//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 메서드 반환 타입을 추출하고 Promise를 언래핑한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractReturnType returns the unwrapped return type of a method.
// If the return type is Promise<X>, it returns X.
func extractReturnType(m *sitter.Node, src []byte) string {
	ann := findChildByType(m, "type_annotation")
	if ann == nil {
		return ""
	}
	raw := extractTypeAnnotation(ann, src)
	return unwrapPromise(raw)
}
