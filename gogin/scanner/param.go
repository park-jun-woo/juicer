//ff:type feature=scan type=model
//ff:what Param 데이터 구조
package scanner

// Param — 요청 파라미터
type Param struct {
	Name    string `yaml:"name"              json:"name"`
	Type    string `yaml:"type"              json:"type"`
	Default string `yaml:"default,omitempty" json:"default,omitempty"`
}
