//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what HTTP 메서드에 따른 기본 상태 코드를 반환한다
package dotnet

func defaultStatusForMethod(method string) string {
	if method == "POST" {
		return "201"
	}
	return "200"
}
