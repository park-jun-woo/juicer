//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestMergeRoles 테스트
package dotnet

import "testing"

func TestMergeRoles(t *testing.T) {
	if mergeRoles([]string{"a"}, []string{"b"})[0] != "b" {
		t.Fatal("method")
	}
	if mergeRoles([]string{"a"}, nil)[0] != "a" {
		t.Fatal("class")
	}
}
