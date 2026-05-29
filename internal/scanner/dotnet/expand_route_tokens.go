//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what [controller]와 [action] 토큰을 실제 값으로 치환한다
package dotnet

import "strings"

func expandRouteTokens(route, controllerName, actionName string) string {
	route = strings.ReplaceAll(route, "[controller]", resolveControllerName(controllerName))
	route = strings.ReplaceAll(route, "[action]", strings.ToLower(actionName))
	return stripRouteConstraints(route)
}
