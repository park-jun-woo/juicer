//ff:func feature=scan type=parse control=iteration dimension=1 topic=nestjs
//ff:what 데코레이터 호출식에서 이름과 인자를 파싱한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// parseDecoratorCall handles @Name('arg') or @Name(value).
func parseDecoratorCall(call *sitter.Node, src []byte) decoratorInfo {
	d := decoratorInfo{}
	fn := findChildByType(call, "identifier")
	if fn != nil {
		d.name = nodeText(fn, src)
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return d
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		arg := args.Child(i)
		switch arg.Type() {
		case "string", "template_string":
			d.arg = unquoteTS(nodeText(arg, src))
			return d
		case "number":
			d.arg = nodeText(arg, src)
			return d
		case "object":
			d.objectProps = make(map[string]string)
			parseObjectArg(arg, src, &d)
			return d
		}
	}
	return d
}
