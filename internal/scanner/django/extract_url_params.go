//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what Django URL 패턴(path 및 re_path)에서 path parameter를 추출한다
package django

// extractURLParams extracts URL variable definitions from a Django URL pattern.
// e.g. "users/<int:pk>/posts/<slug:slug>" -> [{name:"pk", converter:"int"}, {name:"slug", converter:"slug"}];
// re_path named groups "(?P<year>[0-9]+)" -> [{name:"year"}].
func extractURLParams(path string) []urlParam {
	var params []urlParam
	for _, m := range djangoRePathNamedRe.FindAllStringSubmatch(path, -1) {
		params = append(params, urlParam{name: m[1]})
	}
	// Convert named groups to {name} so djangoParamRe does not double-match the inner <name>.
	rest := djangoRePathNamedRe.ReplaceAllString(path, "{$1}")
	for _, m := range djangoParamRe.FindAllStringSubmatch(rest, -1) {
		params = append(params, urlParam{name: m[2], converter: m[1]})
	}
	return params
}
