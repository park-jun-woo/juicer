//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 경로 템플릿에서 :email, {email} 패턴의 파라미터명을 추출한다
package nestjs

import "strings"

// extractPathParamNames returns parameter names from a route path template.
// It recognises both :name and {name} patterns.
func extractPathParamNames(path string) []string {
	var names []string
	for _, seg := range strings.Split(path, "/") {
		if strings.HasPrefix(seg, ":") {
			names = append(names, seg[1:])
		} else if strings.HasPrefix(seg, "{") && strings.HasSuffix(seg, "}") {
			names = append(names, seg[1:len(seg)-1])
		}
	}
	return names
}
