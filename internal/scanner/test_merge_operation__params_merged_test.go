//ff:func feature=scan type=extract control=sequence
//ff:what TestMergeOperation_ParamsMerged 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeOperation_ParamsMerged(t *testing.T) {
	// scan에 parameter가 있고, base에 같은 이름의 parameter에 description이 있음
	scanParam := &yaml.Node{Kind: yaml.MappingNode}
	scanParam.Content = append(scanParam.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "name"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "id"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "in"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "path"},
	)
	scanParams := &yaml.Node{Kind: yaml.SequenceNode}
	scanParams.Content = append(scanParams.Content, scanParam)

	scanOp := &yaml.Node{Kind: yaml.MappingNode}
	scanOp.Content = append(scanOp.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "parameters"},
		scanParams,
	)

	baseParam := &yaml.Node{Kind: yaml.MappingNode}
	baseParam.Content = append(baseParam.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "name"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "id"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "description"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "The user ID"},
	)
	baseParams := &yaml.Node{Kind: yaml.SequenceNode}
	baseParams.Content = append(baseParams.Content, baseParam)

	baseOp := &yaml.Node{Kind: yaml.MappingNode}
	baseOp.Content = append(baseOp.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "parameters"},
		baseParams,
	)

	merged := mergeOperation(scanOp, baseOp)
	params := findMappingValue(merged, "parameters")
	if params == nil || params.Kind != yaml.SequenceNode || len(params.Content) == 0 {
		t.Fatal("expected merged parameters")
	}

	desc := findMappingValue(params.Content[0], "description")
	if desc == nil || desc.Value != "The user ID" {
		t.Error("expected parameter description preserved from base")
	}
}
