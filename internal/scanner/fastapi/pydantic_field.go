//ff:type feature=scan type=model topic=fastapi
//ff:what Pydantic 모델 필드 구조체
package fastapi

// pydanticField holds a single field from a Pydantic model.
type pydanticField struct {
	name       string
	typeName   string
	hasDefault bool
	nullable   bool
	ge         *int
	le         *int
	minLength  *int
	maxLength  *int
}
