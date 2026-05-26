//ff:type feature=scan type=model
//ff:what Field 데이터 구조
package scanner

// Field — struct 필드 정보
type Field struct {
	Name     string  `yaml:"name"              json:"name"`
	Type     string  `yaml:"type"              json:"type"`
	JSON     string  `yaml:"json,omitempty"    json:"json,omitempty"`
	Validate string  `yaml:"validate,omitempty" json:"validate,omitempty"`
	Nullable bool    `yaml:"nullable,omitempty" json:"nullable,omitempty"`
	Fields   []Field `yaml:"fields,omitempty"  json:"fields,omitempty"`
}
