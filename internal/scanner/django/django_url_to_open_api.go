//ff:func feature=scan type=convert control=sequence topic=django
//ff:what Django URL 패턴의 <type:name> 및 re_path (?P<name>..)를 OpenAPI {name}로 변환한다
package django

import "strings"

// djangoURLToOpenAPI converts Django URL patterns to OpenAPI path format.
// e.g. "users/<int:pk>/" -> "users/{pk}/"; re_path "^articles/(?P<year>[0-9]+)/$" -> "articles/{year}/".
func djangoURLToOpenAPI(path string) string {
	path = djangoRePathNamedRe.ReplaceAllString(path, "{$1}")
	path = djangoParamRe.ReplaceAllString(path, "{$2}")
	path = strings.TrimPrefix(path, "^")
	path = strings.TrimSuffix(path, "$")
	return path
}
