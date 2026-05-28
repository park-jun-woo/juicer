//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what 데코레이터 메서드명으로 HTTP 메서드 목록을 결정한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// resolveHTTPMethods determines HTTP methods from a decorator method name.
// For shortcut methods (@app.get, @app.post), returns the single method.
// For @app.route, extracts from methods= argument (default: GET).
func resolveHTTPMethods(methodName string, args *sitter.Node, src []byte) []string {
	if shortcut, ok := shortcutMethods[methodName]; ok {
		return []string{shortcut}
	}
	if methodName == "route" {
		methods := extractMethodsArg(args, src)
		if len(methods) == 0 {
			return []string{"GET"}
		}
		return methods
	}
	return nil
}
