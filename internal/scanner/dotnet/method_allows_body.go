//ff:func feature=scan type=extract control=selection topic=dotnet
//ff:what HTTP 메서드가 요청 바디를 허용하는지(POST/PUT/PATCH) 확인한다
package dotnet

func methodAllowsBody(method string) bool {
	switch method {
	case "POST", "PUT", "PATCH":
		return true
	default:
		return false
	}
}
