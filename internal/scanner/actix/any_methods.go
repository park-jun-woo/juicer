//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what method 미지정 리소스(.to() 직속)가 매칭하는 HTTP 메서드 목록을 반환한다
package actix

// anyMethods returns the set of HTTP methods a method-less resource handler
// (web::resource("...").to(h)) responds to. Actix registers it for any method;
// following the cross-scanner ALL/ANY convention it expands to these five.
func anyMethods() []string {
	return []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
}
