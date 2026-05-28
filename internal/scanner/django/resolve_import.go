//ff:func feature=scan type=extract control=sequence topic=django
//ff:what import 구문에서 View/Serializer 클래스의 정의 파일을 추적한다
package django

import "strings"

// resolveViewName resolves a dotted view name to just the class/function name.
// e.g., "views.UserViewSet" -> "UserViewSet", "UserViewSet" -> "UserViewSet"
func resolveViewName(name string) string {
	parts := strings.Split(name, ".")
	return parts[len(parts)-1]
}
