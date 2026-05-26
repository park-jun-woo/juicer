//ff:type feature=scan type=model topic=fastapi
//ff:what OpenAPI 타입 변환 결과 구조체
package fastapi

// openAPIType holds OpenAPI type and optional format.
type openAPIType struct {
	Type     string
	Format   string
	Items    string // for array types
	Nullable bool
}
