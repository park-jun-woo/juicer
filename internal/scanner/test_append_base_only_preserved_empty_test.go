//ff:func feature=scan type=test control=sequence
//ff:what TestAppendBaseOnlyPreserved_Empty 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestAppendBaseOnlyPreserved_Empty(t *testing.T) {
	result := &yaml.Node{Kind: yaml.MappingNode}
	baseOp := &yaml.Node{Kind: yaml.MappingNode}
	appendBaseOnlyPreserved(result, baseOp, nil)
	if len(result.Content) != 0 {
		t.Fatal("expected no content")
	}
}
