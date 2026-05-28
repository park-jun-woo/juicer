//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what ViewSet 부모 클래스에서 제공하는 HTTP 메서드를 결정한다
package django

// resolveViewSetMethods resolves the HTTP methods provided by a ViewSet's parent classes.
func resolveViewSetMethods(parents []string) []actionMethod {
	var methods []actionMethod
	seen := make(map[string]bool)
	for _, parent := range parents {
		ms := viewsetMethods[parent]
		methods = appendUnseenMethods(methods, ms, seen)
	}
	return methods
}
