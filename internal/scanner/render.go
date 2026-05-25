//ff:func feature=scan type=extract control=selection
//ff:what 스캔 결과를 YAML 또는 JSON으로 직렬화한다
package scanner

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func Render(result *ScanResult, format Format) ([]byte, error) {
	switch format {
	case FormatYAML:
		return yaml.Marshal(result)
	case FormatJSON:
		return json.MarshalIndent(result, "", "  ")
	case FormatOpenAPI:
		return ToOpenAPI(result)
	default:
		return nil, fmt.Errorf("unknown format: %d", format)
	}
}

