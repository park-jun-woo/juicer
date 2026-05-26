//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what ScanResult에서 OpenAPI 3.0 최상위 yaml.Node를 조립한다 (키 순서 보장)
package scanner

import (
	"strings"

	"gopkg.in/yaml.v3"
)

func buildSpecNode(result *ScanResult) *yaml.Node {
	schemas := map[string]any{}
	paths := map[string]map[string]any{}

	deduplicated := deduplicateEndpoints(result.Endpoints)

	for _, ep := range deduplicated {
		oaPath := ginPathToOpenAPI(ep.Path)
		if paths[oaPath] == nil {
			paths[oaPath] = map[string]any{}
		}

		method := strings.ToLower(ep.Method)
		// Gin "Any" → OpenAPI에 해당하는 단일 메서드 없음, get으로 대표
		if method == "any" {
			method = "get"
		}

		op := buildOperation(ep, schemas)
		paths[oaPath][method] = op
	}

	// 키 순서: openapi → info → paths → components
	root := &yaml.Node{Kind: yaml.MappingNode}

	root.Content = append(root.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "openapi"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "3.0.3", Style: yaml.DoubleQuotedStyle},
	)

	infoNode := &yaml.Node{Kind: yaml.MappingNode}
	infoNode.Content = append(infoNode.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "title"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "API (extracted by juicer)", Style: yaml.DoubleQuotedStyle},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "version"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "0.0.0", Style: yaml.DoubleQuotedStyle},
	)
	root.Content = append(root.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "info"},
		infoNode,
	)

	if len(paths) > 0 {
		root.Content = append(root.Content,
			&yaml.Node{Kind: yaml.ScalarNode, Value: "paths"},
			sortedPaths(paths),
		)
	}

	if len(schemas) > 0 {
		compNode := &yaml.Node{Kind: yaml.MappingNode}
		compNode.Content = append(compNode.Content,
			&yaml.Node{Kind: yaml.ScalarNode, Value: "schemas"},
			toYAMLNode(schemas),
		)
		root.Content = append(root.Content,
			&yaml.Node{Kind: yaml.ScalarNode, Value: "components"},
			compNode,
		)
	}

	return root
}

