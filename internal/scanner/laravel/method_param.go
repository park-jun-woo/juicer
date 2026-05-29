//ff:type feature=scan type=model topic=laravel
//ff:what 컨트롤러 메서드 파라미터(이름/타입명)
package laravel

// methodParam holds a parameter of a controller method.
type methodParam struct {
	name     string
	typeName string
}
