//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what web::scope("/prefix") 스코프에서 prefix를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type scopeInfo struct {
	prefix   string
	handlers []string // handler names registered via .service(handler_name)
}

func extractScopes(fi *fileInfo) []scopeInfo {
	var scopes []scopeInfo
	walkNodes(fi.root, func(n *sitter.Node) {
		if n.Type() != "call_expression" {
			return
		}
		scopedID := findChildByType(n, "scoped_identifier")
		if scopedID == nil {
			return
		}
		if nodeText(scopedID, fi.src) != "web::scope" {
			return
		}
		prefix := extractFirstStringArg(n, fi.src)
		if prefix == "" {
			return
		}

		// Collect .service(handler) calls in the chain
		handlers := collectServiceHandlers(n, fi.src)
		scopes = append(scopes, scopeInfo{
			prefix:   prefix,
			handlers: handlers,
		})
	})
	return scopes
}

func collectServiceHandlers(scopeCallNode *sitter.Node, src []byte) []string {
	var handlers []string
	// Walk up the chain: scope("...").service(X).service(Y)
	// The scope call is wrapped in field_expression -> call_expression chains
	parent := scopeCallNode.Parent()
	for parent != nil {
		if parent.Type() == "call_expression" {
			fe := findChildByType(parent, "field_expression")
			if fe != nil {
				fid := findChildByType(fe, "field_identifier")
				if fid != nil && nodeText(fid, src) == "service" {
					args := findChildByType(parent, "arguments")
					if args != nil {
						for i := 0; i < int(args.ChildCount()); i++ {
							arg := args.Child(i)
							if arg.Type() == "identifier" {
								handlers = append(handlers, nodeText(arg, src))
							}
						}
					}
				}
			}
		}
		parent = parent.Parent()
	}
	return handlers
}
