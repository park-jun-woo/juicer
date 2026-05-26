//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what path item 내의 operation 중 등록된 것만 필터링한다
package scanner

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func filterOperations(ops *yaml.Node, path string, registered map[string]bool) (*yaml.Node, bool) {
	if ops.Kind != yaml.MappingNode {
		return ops, false
	}
	filtered := &yaml.Node{Kind: yaml.MappingNode}
	hasAny := false
	for i := 0; i+1 < len(ops.Content); i += 2 {
		method := ops.Content[i].Value
		key := method + "\t" + path
		if registered[key] {
			filtered.Content = append(filtered.Content, ops.Content[i], ops.Content[i+1])
			hasAny = true
			continue
		}
		fmt.Fprintf(os.Stderr, "warning: dead spec dropped: %s %s\n", strings.ToUpper(method), path)
	}
	return filtered, hasAny
}
