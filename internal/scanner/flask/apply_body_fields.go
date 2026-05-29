//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 추출한 바디 필드를 각 routeInfo에 주입한다
package flask

// applyBodyFields injects collected body fields into every routeInfo produced
// for a handler (all methods of the same handler share its body access).
func applyBodyFields(routes []routeInfo, bf bodyFields) []routeInfo {
	for i := range routes {
		routes[i].formFields = bf.formFields
		routes[i].jsonFields = bf.jsonFields
		routes[i].hasJSONBody = bf.hasJSON
	}
	return routes
}
