//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what register(autoload, {dir, options:{prefix}}) 호출에서 dir 경로와 베이스 prefix를 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractAutoloadCall(call *sitter.Node, src []byte, names map[string]bool) (string, string, bool) {
	fn := findChildByType(call, "member_expression")
	if fn == nil {
		return "", "", false
	}
	prop := fn.ChildByFieldName("property")
	if prop == nil || nodeText(prop, src) != "register" {
		return "", "", false
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return "", "", false
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) < 2 || !isAutoloadArg(argNodes[0], src, names) {
		return "", "", false
	}
	dir := extractAutoloadDir(argNodes[1], src)
	if dir == "" {
		return "", "", false
	}
	return dir, extractAutoloadPrefix(argNodes[1], src), true
}
