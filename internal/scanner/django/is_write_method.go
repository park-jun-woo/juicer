//ff:func feature=scan type=extract control=sequence topic=django
//ff:what HTTP 메서드가 쓰기 메서드(POST/PUT/PATCH)인지 확인한다
package django

// isWriteMethod returns true for HTTP methods that typically carry a request body.
func isWriteMethod(method string) bool {
	return method == "POST" || method == "PUT" || method == "PATCH"
}
