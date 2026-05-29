//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 단일 call_expression에서 fastify.register() 호출을 파싱한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractRegisterCall(call *sitter.Node, src []byte, instances map[string]bool) *pluginMount {
	fn := findChildByType(call, "member_expression")
	if fn == nil {
		return nil
	}
	obj := findChildByType(fn, "identifier")
	if obj == nil || !instances[nodeText(obj, src)] {
		return nil
	}
	prop := fn.ChildByFieldName("property")
	if prop == nil || nodeText(prop, src) != "register" {
		return nil
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) == 0 {
		return nil
	}
	pm := &pluginMount{PluginRef: extractPluginRef(argNodes[0], src)}
	if len(argNodes) >= 2 && argNodes[1].Type() == "object" {
		pm.Prefix = extractPrefixFromOpts(argNodes[1], src)
	}
	if pm.PluginRef == inlineRef {
		pm.Inline = true
		pm.WrapperStart = argNodes[0].StartByte()
		pm.WrapperEnd = argNodes[0].EndByte()
	}
	return pm
}
