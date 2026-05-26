//ff:func feature=scan type=parse control=iteration dimension=1 topic=nestjs
//ff:what 데코레이터 호출식에서 이름과 인자를 파싱한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// parseDecoratorCall handles @Name('arg') or @Name(value).
// It supports string, template_string, number, object, identifier, and
// member_expression argument types. All non-object arguments are collected
// into d.args; d.arg is set to the first argument for backward compatibility.
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
			v := unquoteTS(nodeText(arg, src))
			d.args = append(d.args, v)
		case "number":
			d.args = append(d.args, nodeText(arg, src))
		case "identifier":
			d.args = append(d.args, nodeText(arg, src))
		case "member_expression":
			d.args = append(d.args, nodeText(arg, src))
		case "object":
			d.objectProps = make(map[string]string)
			parseObjectArg(arg, src, &d)
			return d
		}
	}
	if len(d.args) > 0 {
		d.arg = d.args[0]
	}
	return d
}
