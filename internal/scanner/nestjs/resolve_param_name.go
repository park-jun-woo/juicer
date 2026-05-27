//ff:func feature=scan type=extract control=selection topic=nestjs
//ff:what 데코레이터 arg가 빈 문자열일 때 파라미터 이름을 결정한다
package nestjs

// resolveParamName returns the effective parameter name for a decorator.
// When the decorator argument is non-empty it is used as-is.
// For an empty @Param() argument, a single-segment route path yields the path
// parameter name; otherwise the Go parameter name is returned as fallback.
func resolveParamName(decName, arg, paramName, routePath string) string {
	if arg != "" {
		return arg
	}
	switch decName {
	case DecParam:
		names := extractPathParamNames(routePath)
		if len(names) == 1 {
			return names[0]
		}
		return paramName
	default:
		return paramName
	}
}
