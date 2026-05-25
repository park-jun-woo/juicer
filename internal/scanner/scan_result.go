//ff:type feature=scan type=model
//ff:what ScanResult 데이터 구조
package scanner

// ScanResult — 스캔 결과 최상위 구조
type ScanResult struct {
	Endpoints []Endpoint `yaml:"endpoints" json:"endpoints"`
}
