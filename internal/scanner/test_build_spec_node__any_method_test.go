//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestBuildSpecNode_AnyMethod — Any 메서드가 5개 HTTP 메서드로 등록되는지 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestBuildSpecNode_AnyMethod(t *testing.T) {
	result := &ScanResult{
		Endpoints: []Endpoint{
			{Method: "Any", Path: "/health", Handler: "main.go:health"},
		},
	}
	node := buildSpecNode(result)
	if node == nil {
		t.Fatal("expected non-nil")
	}

	// paths 노드를 찾는다
	var pathsNode *yaml.Node
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == "paths" {
			pathsNode = node.Content[i+1]
			break
		}
	}
	if pathsNode == nil {
		t.Fatal("expected paths node")
	}

	// /health path item을 찾는다
	var healthOps *yaml.Node
	for i := 0; i+1 < len(pathsNode.Content); i += 2 {
		if pathsNode.Content[i].Value == "/health" {
			healthOps = pathsNode.Content[i+1]
			break
		}
	}
	if healthOps == nil {
		t.Fatal("expected /health path item")
	}

	// 5개 메서드가 있는지 확인
	methods := map[string]bool{}
	for i := 0; i+1 < len(healthOps.Content); i += 2 {
		methods[healthOps.Content[i].Value] = true
	}
	for _, m := range []string{"get", "post", "put", "patch", "delete"} {
		if !methods[m] {
			t.Errorf("expected method %s for /health, got methods: %v", m, methods)
		}
	}
}
