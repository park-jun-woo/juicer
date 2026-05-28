//ff:func feature=scan type=convert control=sequence topic=django
//ff:what Django URL 패턴의 <type:name>을 OpenAPI {name} 형식으로 변환한다
package django

// djangoURLToOpenAPI converts Django URL patterns to OpenAPI path format.
// e.g., "users/<int:pk>/" -> "users/{pk}/"
func djangoURLToOpenAPI(path string) string {
	return djangoParamRe.ReplaceAllString(path, "{$2}")
}
