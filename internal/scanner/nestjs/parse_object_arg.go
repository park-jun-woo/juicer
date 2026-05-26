//ff:func feature=scan type=parse control=iteration dimension=1 topic=nestjs
//ff:what 객체 리터럴 인자에서 key-value 쌍을 파싱한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// parseObjectArg extracts key-value pairs from an object literal argument.
// For @Controller({ path: 'auth', version: '1' }), it sets d.arg to the path
// value and stores all string properties in d.objectProps.
func parseObjectArg(obj *sitter.Node, src []byte, d *decoratorInfo) {
	for i := 0; i < int(obj.ChildCount()); i++ {
		child := obj.Child(i)
		if child.Type() != "pair" {
			continue
		}
		key := findChildByType(child, "property_identifier")
		if key == nil {
			continue
		}
		keyText := nodeText(key, src)
		val := findPairStringValue(child, src)
		if val == "" {
			continue
		}
		d.objectProps[keyText] = val
		if keyText == "path" {
			d.arg = val
		}
	}
}
