//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 팩토리 함수 내부에 선언된 클래스 노드를 찾는다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// findFactoryInnerClass locates a function_declaration with the given name
// and returns the first class_declaration found inside its statement_block.
func findFactoryInnerClass(root *sitter.Node, src []byte, funcName string) *sitter.Node {
	funcs := findAllByType(root, "function_declaration")
	for _, fn := range funcs {
		nameNode := findChildByType(fn, "identifier")
		if nameNode == nil || nodeText(nameNode, src) != funcName {
			continue
		}
		body := findChildByType(fn, "statement_block")
		if body == nil {
			continue
		}
		classes := findAllByType(body, "class_declaration")
		if len(classes) > 0 {
			return classes[0]
		}
	}
	return nil
}
