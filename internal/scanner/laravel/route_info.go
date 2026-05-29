//ff:type feature=scan type=model topic=laravel
//ff:what 추출된 단일 Laravel 라우트 정보
package laravel

// routeInfo holds a single extracted route.
type routeInfo struct {
	method     string // HTTP method (uppercase)
	path       string // URL path
	controller string // controller class name
	action     string // controller method name
	file       string // source file relative path
	line       int    // source line number
	middleware []string
}
