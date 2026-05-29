//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what invocation_expression에서 MapGet/MapPost 등을 매칭한다
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func matchMapMethod(inv *sitter.Node, fi *fileInfo, groups map[string]string) (scanner.Endpoint, bool) {
	receiver, methodName := extractMethodCall(inv, fi.src)
	httpMethod, ok := mapMethods[methodName]
	if !ok {
		return scanner.Endpoint{}, false
	}

	args := findChildByType(inv, "argument_list")
	if args == nil {
		return scanner.Endpoint{}, false
	}
	argNodes := childrenOfType(args, "argument")
	if len(argNodes) < 1 {
		return scanner.Endpoint{}, false
	}

	pathLit := findChildByType(argNodes[0], "string_literal")
	if pathLit == nil {
		return scanner.Endpoint{}, false
	}
	path := stripRouteConstraints(unquoteCSharp(nodeText(pathLit, fi.src)))

	prefix := ""
	if p, ok := groups[receiver]; ok {
		prefix = p
	}
	fullPath := joinPath(prefix, path)

	ep := scanner.Endpoint{
		Method:  httpMethod,
		Path:    fullPath,
		Handler: methodName,
		File:    fi.relPath,
		Line:    int(inv.StartPoint().Row) + 1,
	}

	ep.Request = extractLambdaRequest(argNodes, fi.src, path)

	return ep, true
}
