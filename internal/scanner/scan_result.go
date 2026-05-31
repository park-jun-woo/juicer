//ff:type feature=scan type=model
//ff:what ScanResult 데이터 구조
package scanner

// ScanResult — 스캔 결과 최상위 구조
type ScanResult struct {
	Endpoints []Endpoint `yaml:"endpoints" json:"endpoints"`
	// Schemas holds extra named component schemas discovered during scanning
	// (e.g. nested DTOs/enums recursively extracted by the NestJS scanner).
	// Each value is an OpenAPI schema object. Keyed by component schema name.
	Schemas map[string]any `yaml:"schemas,omitempty" json:"schemas,omitempty"`
}
