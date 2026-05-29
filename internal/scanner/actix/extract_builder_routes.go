//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 빌더 패턴(web::resource().route(), web::scope().service())에서 라우트를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type builderRoute struct {
	method  string
	path    string
	handler string
}

func extractBuilderRoutes(fi *fileInfo) []builderRoute {
	var routes []builderRoute
	// Find top-level .service() calls (on App, cfg, etc.) — not nested ones
	walkNodes(fi.root, func(n *sitter.Node) {
		if n.Type() != "call_expression" {
			return
		}
		fe := findChildByType(n, "field_expression")
		if fe == nil {
			return
		}
		fid := findChildByType(fe, "field_identifier")
		if fid == nil {
			return
		}
		if nodeText(fid, fi.src) != "service" {
			return
		}
		// Check that the receiver is NOT a web::scope or web::resource chain
		// (those are handled recursively by processServiceArg)
		receiver := findFieldReceiver(fe)
		if receiver != nil && isWebScopeOrResource(receiver, fi.src) {
			return
		}
		args := findChildByType(n, "arguments")
		if args == nil {
			return
		}
		for i := 0; i < int(args.ChildCount()); i++ {
			arg := args.Child(i)
			if arg.Type() == "call_expression" {
				processServiceArg(arg, fi.src, "", &routes)
			}
		}
	})
	return deduplicateBuilderRoutes(routes)
}

func isWebScopeOrResource(node *sitter.Node, src []byte) bool {
	root := findCallRoot(node, src)
	return root == "web::scope" || root == "web::resource"
}

func findFieldReceiver(fieldExpr *sitter.Node) *sitter.Node {
	// The receiver is the first child of field_expression
	if fieldExpr.ChildCount() > 0 {
		return fieldExpr.Child(0)
	}
	return nil
}

func processServiceArg(callExpr *sitter.Node, src []byte, prefix string, routes *[]builderRoute) {
	root := findCallRoot(callExpr, src)
	switch root {
	case "web::scope":
		scopePrefix := extractScopePrefix(callExpr, src)
		collectServiceCalls(callExpr, src, joinPath(prefix, scopePrefix), routes)
	case "web::resource":
		resourcePath := extractResourcePath(callExpr, src)
		collectRouteCalls(callExpr, src, joinPath(prefix, resourcePath), routes)
	}
}

func collectServiceCalls(node *sitter.Node, src []byte, prefix string, routes *[]builderRoute) {
	// Find .service() calls directly chained on this scope
	findServiceCalls(node, src, func(args *sitter.Node) {
		for i := 0; i < int(args.ChildCount()); i++ {
			arg := args.Child(i)
			if arg.Type() == "call_expression" {
				processServiceArg(arg, src, prefix, routes)
			}
		}
	})
}

func collectRouteCalls(node *sitter.Node, src []byte, resourcePath string, routes *[]builderRoute) {
	// Find .route() calls directly chained on this resource
	findRouteCalls(node, src, func(args *sitter.Node) {
		method, handler := parseRouteArg(args, src)
		if method != "" {
			*routes = append(*routes, builderRoute{
				method:  method,
				path:    resourcePath,
				handler: handler,
			})
		}
	})
}

func findServiceCalls(node *sitter.Node, src []byte, fn func(*sitter.Node)) {
	walkMethodChain(node, src, "service", fn)
}

func findRouteCalls(node *sitter.Node, src []byte, fn func(*sitter.Node)) {
	walkMethodChain(node, src, "route", fn)
}

// walkMethodChain descends a method-call chain from its outermost call toward
// the chain head, invoking fn(arguments) for each .<method>() call on the way.
// Pattern: web::scope("...").service(X).service(Y) nests with the outermost
// call (.service(Y)) wrapping its receiver (.service(X)), down to the head
// (web::scope("...")). Walking DOWN the receiver chain — rather than up through
// ancestors — keeps the traversal inside this chain and avoids re-entering the
// enclosing .service() that wraps it (which caused unbounded recursion).
func walkMethodChain(node *sitter.Node, src []byte, method string, fn func(*sitter.Node)) {
	for n := node; n != nil && n.Type() == "call_expression"; {
		fe := findChildByType(n, "field_expression")
		if fe == nil {
			// Reached the chain head, e.g. web::scope("...") whose callee is a
			// scoped_identifier rather than a field_expression.
			return
		}
		fid := findChildByType(fe, "field_identifier")
		if fid != nil && nodeText(fid, src) == method {
			if args := findChildByType(n, "arguments"); args != nil {
				fn(args)
			}
		}
		// Descend into the receiver (first child of the field_expression).
		n = fe.Child(0)
	}
}

func parseRouteArg(args *sitter.Node, src []byte) (string, string) {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "call_expression" {
			continue
		}
		fe := findChildByType(child, "field_expression")
		if fe == nil {
			continue
		}
		fid := findChildByType(fe, "field_identifier")
		if fid == nil || nodeText(fid, src) != "to" {
			continue
		}
		method := extractWebMethod(fe, src)
		handler := extractToHandler(child, src)
		return method, handler
	}
	return "", ""
}

func extractWebMethod(fieldExpr *sitter.Node, src []byte) string {
	for i := 0; i < int(fieldExpr.ChildCount()); i++ {
		child := fieldExpr.Child(i)
		if child.Type() == "call_expression" {
			scopedID := findChildByType(child, "scoped_identifier")
			if scopedID != nil {
				parts := splitScoped(nodeText(scopedID, src))
				if len(parts) == 2 && parts[0] == "web" {
					if m, ok := webMethodBuilders[parts[1]]; ok {
						return m
					}
				}
			}
		}
	}
	return ""
}

func extractToHandler(callExpr *sitter.Node, src []byte) string {
	args := findChildByType(callExpr, "arguments")
	if args == nil {
		return ""
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "identifier" {
			return nodeText(child, src)
		}
	}
	return ""
}

func findCallRoot(node *sitter.Node, src []byte) string {
	var result string
	walkNodes(node, func(n *sitter.Node) {
		if n.Type() == "scoped_identifier" && result == "" {
			text := nodeText(n, src)
			if text == "web::scope" || text == "web::resource" {
				result = text
			}
		}
	})
	return result
}

func extractScopePrefix(node *sitter.Node, src []byte) string {
	var result string
	walkNodes(node, func(n *sitter.Node) {
		if result != "" {
			return
		}
		if n.Type() == "scoped_identifier" && nodeText(n, src) == "web::scope" {
			parent := n.Parent()
			if parent != nil && parent.Type() == "call_expression" {
				result = extractFirstStringArg(parent, src)
			}
		}
	})
	return result
}

func extractResourcePath(node *sitter.Node, src []byte) string {
	var result string
	walkNodes(node, func(n *sitter.Node) {
		if result != "" {
			return
		}
		if n.Type() == "scoped_identifier" && nodeText(n, src) == "web::resource" {
			parent := n.Parent()
			if parent != nil && parent.Type() == "call_expression" {
				result = extractFirstStringArg(parent, src)
			}
		}
	})
	return result
}

func extractFirstStringArg(callExpr *sitter.Node, src []byte) string {
	args := findChildByType(callExpr, "arguments")
	if args == nil {
		return ""
	}
	strLit := findChildByType(args, "string_literal")
	if strLit == nil {
		return ""
	}
	strContent := findChildByType(strLit, "string_content")
	if strContent == nil {
		return ""
	}
	return nodeText(strContent, src)
}

func deduplicateBuilderRoutes(routes []builderRoute) []builderRoute {
	seen := map[string]bool{}
	var result []builderRoute
	for _, r := range routes {
		key := r.method + " " + r.path + " " + r.handler
		if seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, r)
	}
	return result
}
