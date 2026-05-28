//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what Django URL 패턴에서 path parameter를 추출한다
package django

// extractURLParams extracts URL variable definitions from a Django URL pattern.
// e.g., "users/<int:pk>/posts/<slug:slug>" -> [{name:"pk", converter:"int"}, {name:"slug", converter:"slug"}]
func extractURLParams(path string) []urlParam {
	matches := djangoParamRe.FindAllStringSubmatch(path, -1)
	var params []urlParam
	for _, m := range matches {
		params = append(params, urlParam{
			name:      m[2],
			converter: m[1],
		})
	}
	return params
}
