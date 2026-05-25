//ff:type feature=scan type=model
//ff:what Request 데이터 구조
package scanner

// Request — 엔드포인트 요청 정보
type Request struct {
	PathParams []Param `yaml:"path_params,omitempty" json:"path_params,omitempty"`
	Body       *Body   `yaml:"body,omitempty"        json:"body,omitempty"`
	Query      []Param `yaml:"query,omitempty"       json:"query,omitempty"`
	FormFields []Param `yaml:"form_fields,omitempty" json:"form_fields,omitempty"`
	Files      []Param `yaml:"files,omitempty"       json:"files,omitempty"`
	RawBody    bool    `yaml:"raw_body,omitempty"    json:"raw_body,omitempty"`
}
