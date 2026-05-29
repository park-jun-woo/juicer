//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 파일에서 top-level `const <name> = { ... }` 선언의 object value 노드를 찾는다
package express

import sitter "github.com/smacker/go-tree-sitter"

// findConstObject — 이름이 name인 변수 선언의 object 리터럴 value 노드를 반환한다.
func findConstObject(root *sitter.Node, src []byte, name string) *sitter.Node {
	for _, decl := range findAllByType(root, "lexical_declaration") {
		if obj := constObjectInDecl(decl, src, name); obj != nil {
			return obj
		}
	}
	return nil
}
