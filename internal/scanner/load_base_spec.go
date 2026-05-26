//ff:func feature=scan type=extract control=sequence
//ff:what 기존 openapi.yaml을 yaml.Node로 파싱한다
package scanner

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadBaseSpec(path string) (*yaml.Node, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading base spec: %w", err)
	}
	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("parsing base spec: %w", err)
	}
	if doc.Kind != yaml.DocumentNode || len(doc.Content) == 0 {
		return nil, fmt.Errorf("invalid base spec: empty document")
	}
	return doc.Content[0], nil
}
