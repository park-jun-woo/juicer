//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 메서드 본문의 object_creation_expression에서 new ResponseEntity 상태 코드를 찾는다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func matchBodyConstructors(body *sitter.Node, src []byte) string {
	for _, obj := range findAllByType(body, "object_creation_expression") {
		code := matchResponseEntityConstructor(obj, src)
		if code != "" {
			return code
		}
	}
	return ""
}
