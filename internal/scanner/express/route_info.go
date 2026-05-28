//ff:type feature=scan type=model topic=express
//ff:what 추출된 라우트 정보 구조체
package express

type routeInfo struct {
	Method     string
	Path       string
	Handler    string
	Middleware []string
	Line       int
}
