//ff:func feature=scan type=extract control=sequence
//ff:what ScanResult를 OpenAPI 3.0 spec YAML 바이트로 변환한다
package scanner

import "gopkg.in/yaml.v3"

// OpenAPI 3.0 구조체 — yaml 직렬화용

func ToOpenAPI(result *ScanResult) ([]byte, error) {
	node := buildSpecNode(result)
	return yaml.Marshal(node)
}

