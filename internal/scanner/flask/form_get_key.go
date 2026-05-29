//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what request.form.get('key') 호출에서 폼 필드 키를 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// formGetKey returns the form field key for a `request.form.get('key', ...)`
// call node, or "" if the call is not such an access.
func formGetKey(call *sitter.Node, src []byte) string {
	fn := findChildByType(call, "attribute")
	if fn == nil || nodeText(fn, src) != "request.form.get" {
		return ""
	}
	args := findChildByType(call, "argument_list")
	if args == nil {
		return ""
	}
	return firstStringArg(args, src)
}
