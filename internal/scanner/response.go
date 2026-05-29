//ff:type feature=scan type=model
//ff:what Response 데이터 구조
package scanner

// Response — 응답 정보
type Response struct {
	Status      string  `yaml:"status"                json:"status"`
	Kind        string  `yaml:"kind"                  json:"kind"`
	Body        string  `yaml:"body,omitempty"         json:"body,omitempty"`
	Source      string  `yaml:"source,omitempty"       json:"source,omitempty"`
	TypeName    string  `yaml:"type,omitempty"         json:"type,omitempty"`
	Fields      []Field `yaml:"fields,omitempty"       json:"fields,omitempty"`
	Confidence  string  `yaml:"confidence,omitempty"   json:"confidence,omitempty"`
	ContentType string  `yaml:"contentType,omitempty"  json:"contentType,omitempty"`
}
