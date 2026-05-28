//ff:type feature=scan type=model topic=django
//ff:what OpenAPI 타입 변환 결과 구조체
package django

// openAPIType holds OpenAPI type and optional format.
type openAPIType struct {
	Type   string
	Format string
}
