//ff:type feature=scan type=model topic=actix
//ff:what 빌더 패턴으로 추출한 라우트(method/path/handler)
package actix

type builderRoute struct {
	method  string
	path    string
	handler string
}
