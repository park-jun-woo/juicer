//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 파라미터 이름이 경로 변수에 매칭되는지 확인한다
package dotnet

import "strings"

func isPathParam(name, path string) bool {
	return strings.Contains(path, "{"+name+"}")
}
