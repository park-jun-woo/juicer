//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what 경로 템플릿의 {token} 중 in:path로 선언되지 않은 것을 string path param으로 보강한다
package scanner

func ensurePathParams(params []map[string]any, path string) []map[string]any {
	declared := declaredPathParams(params)
	for _, name := range pathTemplateNames(path) {
		if declared[name] {
			continue
		}
		declared[name] = true
		params = append(params, map[string]any{
			"name":     name,
			"in":       "path",
			"required": true,
			"schema":   buildParamSchema("string"),
		})
	}
	return params
}
