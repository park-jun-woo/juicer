//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what FormRequest 타입 힌트 파라미터의 타입명을 찾는다
package laravel

// findFormRequestParam finds a FormRequest type-hinted parameter.
func findFormRequestParam(params []methodParam) string {
	for _, p := range params {
		if isFormRequestType(p.typeName) {
			return p.typeName
		}
	}
	return ""
}
