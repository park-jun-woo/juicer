//ff:type feature=scan type=model
//ff:what Field 데이터 구조
package scanner

// Field — struct 필드 정보
type Field struct {
	Name      string   `yaml:"name"              json:"name"`
	Type      string   `yaml:"type"              json:"type"`
	JSON      string   `yaml:"json,omitempty"    json:"json,omitempty"`
	Validate  string   `yaml:"validate,omitempty" json:"validate,omitempty"`
	Nullable  bool     `yaml:"nullable,omitempty" json:"nullable,omitempty"`
	Enum      []string `yaml:"enum,omitempty"     json:"enum,omitempty"`
	Minimum   *int     `yaml:"minimum,omitempty"   json:"minimum,omitempty"`
	Maximum   *int     `yaml:"maximum,omitempty"   json:"maximum,omitempty"`
	MinLength *int     `yaml:"min_length,omitempty" json:"min_length,omitempty"`
	MaxLength *int     `yaml:"max_length,omitempty" json:"max_length,omitempty"`
	Fields    []Field  `yaml:"fields,omitempty"  json:"fields,omitempty"`
	// Ref names a component schema this field references (a nested named
	// DTO/enum). When set, the field is emitted as a $ref (array-wrapped when
	// Type=="array") instead of an inline object. The schema itself is expected
	// to be registered separately (see ScanResult.Schemas).
	Ref string `yaml:"ref,omitempty" json:"ref,omitempty"`
}
