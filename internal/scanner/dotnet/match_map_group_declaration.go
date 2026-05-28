//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 로컬 선언에서 MapGroup 호출을 매칭하여 그룹에 추가한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func matchMapGroupDeclaration(stmt *sitter.Node, fi *fileInfo, groups map[string]string) {
	varDecl := findChildByType(stmt, "variable_declaration")
	if varDecl == nil {
		return
	}
	declarator := findChildByType(varDecl, "variable_declarator")
	if declarator == nil {
		return
	}
	nameNode := findChildByType(declarator, "identifier")
	if nameNode == nil {
		return
	}
	varName := nodeText(nameNode, fi.src)

	inv := findChildByType(declarator, "invocation_expression")
	if inv == nil {
		return
	}
	receiver, methodName := extractMethodCall(inv, fi.src)
	if methodName != "MapGroup" {
		return
	}

	args := findChildByType(inv, "argument_list")
	if args == nil {
		return
	}
	prefix := extractFirstStringFromArgs(args, fi.src)
	if prefix == "" {
		return
	}

	parentPrefix := ""
	if p, ok := groups[receiver]; ok {
		parentPrefix = p
	}
	groups[varName] = joinPath(parentPrefix, prefix)
}
