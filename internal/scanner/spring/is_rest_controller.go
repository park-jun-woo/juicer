//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 클래스가 @RestController인지 확인한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func isRestController(cls *sitter.Node, src []byte) bool {
	if hasAnnotation(cls, src, AnnRestController) {
		return true
	}
	if hasAnnotation(cls, src, AnnController) && hasAnnotation(cls, src, AnnResponseBody) {
		return true
	}
	return false
}
