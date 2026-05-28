//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what Hono 경로에서 :param 이름을 추출한다
package hono

import "strings"

func extractPathParams(path string) []string {
	var params []string
	for _, part := range strings.Split(path, "/") {
		name := extractParamName(part)
		if name != "" {
			params = append(params, name)
		}
	}
	return params
}
