//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what ScanResult에서 OpenAPI 3.0 최상위 yaml.Node를 조립한다 (키 순서 보장)
package scanner

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func buildSpecNode(result *ScanResult) *yaml.Node {
	schemas := map[string]any{}
	paths := map[string]map[string]any{}

	deduplicated := deduplicateEndpoints(result.Endpoints)
	confirmedIDs := deduplicateOperationIDs(deduplicated)

	for i, ep := range deduplicated {
		oaPath := ep.Path
		if paths[oaPath] == nil {
			paths[oaPath] = map[string]any{}
		}

		method := strings.ToLower(ep.Method)
		op := buildOperation(ep, schemas)
		if cid, ok := confirmedIDs[i]; ok {
			op["operationId"] = cid
		}
		for _, m := range expandAnyMethod(method) {
			if _, dup := paths[oaPath][m]; dup {
				// 같은 path+method로 붕괴한 충돌. 조용히 덮어쓰지 않고 경고한다.
				fmt.Fprintf(os.Stderr, "warning: duplicate operation %s %s (handler %q at %s:%d) overwrites a previous one — check route prefix composition\n",
					strings.ToUpper(m), oaPath, ep.Handler, ep.File, ep.Line)
			}
			paths[oaPath][m] = op
		}
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
		&yaml.Node{Kind: yaml.ScalarNode, Value: "API (extracted by codist)", Style: yaml.DoubleQuotedStyle},
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

	hasSecurity := false
	for _, ep := range deduplicated {
		if isAuthEndpoint(ep) {
			hasSecurity = true
			break
		}
	}

	if len(schemas) > 0 || hasSecurity {
		compNode := &yaml.Node{Kind: yaml.MappingNode}
		if len(schemas) > 0 {
			compNode.Content = append(compNode.Content,
				&yaml.Node{Kind: yaml.ScalarNode, Value: "schemas"},
				toYAMLNode(schemas),
			)
		}
		if hasSecurity {
			secSchemes := map[string]any{
				"bearerAuth": map[string]any{
					"type":         "http",
					"scheme":       "bearer",
					"bearerFormat": "JWT",
				},
			}
			compNode.Content = append(compNode.Content,
				&yaml.Node{Kind: yaml.ScalarNode, Value: "securitySchemes"},
				toYAMLNode(secSchemes),
			)
		}
		root.Content = append(root.Content,
			&yaml.Node{Kind: yaml.ScalarNode, Value: "components"},
			compNode,
		)
	}

	return root
}
