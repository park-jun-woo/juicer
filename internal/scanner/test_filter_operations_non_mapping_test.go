//ff:func feature=scan type=test control=sequence
//ff:what TestFilterOperations_NonMapping 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFilterOperations_NonMapping(t *testing.T) {
	ops := &yaml.Node{Kind: yaml.ScalarNode, Value: "test"}
	result, hasAny := filterOperations(ops, "/test", nil)
	if hasAny {
		t.Fatal("expected no match for non-mapping node")
	}
	if result != ops {
		t.Fatal("expected original node returned")
	}
}
