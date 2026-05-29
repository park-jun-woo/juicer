//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what 라우트 파라미터 토큰 {name:constraint}/{name=default}/{name?}를 {name}으로 정규화한다
package dotnet

import "regexp"

var routeConstraintRe = regexp.MustCompile(`\{([^}:=?*]+)[^}]*\}`)

func stripRouteConstraints(route string) string {
	return routeConstraintRe.ReplaceAllString(route, "{$1}")
}
