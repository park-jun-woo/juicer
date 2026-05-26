//ff:func feature=scan type=convert control=sequence topic=fastapi
//ff:what HTTP 메서드에 대한 기본 상태코드를 반환한다
package fastapi

// defaultStatusForMethod returns the default HTTP status code for a method.
func defaultStatusForMethod(method string) string {
	return "200"
}
