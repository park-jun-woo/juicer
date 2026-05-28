//ff:type feature=scan type=model topic=django
//ff:what Django URL 변수 파싱 결과 구조체
package django

// urlParam holds a parsed Django URL variable.
type urlParam struct {
	name      string
	converter string // "int", "str", "uuid", "slug", "path", or ""
}
