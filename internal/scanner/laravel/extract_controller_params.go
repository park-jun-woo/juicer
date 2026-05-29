//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 컨트롤러 메서드 파라미터에서 path parameter 타입 힌트를 추출한다
package laravel

import "github.com/park-jun-woo/codistill/internal/scanner"

// applyControllerParamTypes updates path parameter types using controller method type hints.
func applyControllerParamTypes(pathParams []scanner.Param, cm *controllerMethod) []scanner.Param {
	if cm == nil {
		return pathParams
	}
	typeMap := buildParamTypeMap(cm)
	for i := range pathParams {
		if t, ok := typeMap[pathParams[i].Name]; ok {
			pathParams[i].Type = t
		}
	}
	return pathParams
}
