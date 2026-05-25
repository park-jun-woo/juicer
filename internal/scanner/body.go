//ff:type feature=scan type=model
//ff:what Body 데이터 구조
package scanner

// Body — 요청 body binding 정보
type Body struct {
	VarName  string  `yaml:"var_name"          json:"var_name"`
	Method   string  `yaml:"method"            json:"method"`
	TypeName string  `yaml:"type,omitempty"    json:"type,omitempty"`
	Fields   []Field `yaml:"fields,omitempty"  json:"fields,omitempty"`
}
