//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what OpenAPI parameters에서 이미 선언된 in:path 파라미터 이름 집합을 만든다
package scanner

func declaredPathParams(params []map[string]any) map[string]bool {
	declared := make(map[string]bool)
	for _, p := range params {
		name, ok := p["name"].(string)
		if ok && p["in"] == "path" {
			declared[name] = true
		}
	}
	return declared
}
