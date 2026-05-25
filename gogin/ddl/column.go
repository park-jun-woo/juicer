//ff:type feature=ddl type=model
//ff:what 컬럼 정의 (이름 + 원본 텍스트)
package ddl

// Column represents a single column definition.
type Column struct {
	Name string
	Raw  string
}
