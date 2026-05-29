//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what apiResource 액션 하나를 routeInfo로 변환한다(파라미터 suffix 치환 포함)
package laravel

import "fmt"

func apiResourceRoute(action apiResourceAction, fullBase, paramName, controller, relPath string, line int, mw []string) routeInfo {
	suffix := action.suffix
	if action.hasParam {
		suffix = fmt.Sprintf(suffix, paramName)
	}
	return routeInfo{
		method:     action.method,
		path:       fullBase + suffix,
		controller: controller,
		action:     action.action,
		file:       relPath,
		line:       line,
		middleware: mw,
	}
}
