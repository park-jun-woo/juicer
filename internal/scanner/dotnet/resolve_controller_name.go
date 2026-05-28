//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what 클래스명에서 Controller 접미사를 제거하고 소문자로 변환한다
package dotnet

import "strings"

func resolveControllerName(className string) string {
	name := strings.TrimSuffix(className, "Controller")
	return strings.ToLower(name)
}
