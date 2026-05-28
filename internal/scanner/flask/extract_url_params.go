//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what Flask URL 규칙에서 path parameter를 추출한다
package flask

// extractURLParams extracts URL variable definitions from a Flask URL rule.
// e.g., "/users/<int:user_id>/posts/<post_id>" -> [{name:"user_id", converter:"int"}, {name:"post_id", converter:""}]
func extractURLParams(path string) []urlParam {
	matches := flaskParamRe.FindAllStringSubmatch(path, -1)
	var params []urlParam
	for _, m := range matches {
		params = append(params, urlParam{
			name:      m[2],
			converter: m[1],
		})
	}
	return params
}
