//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractMapGroups_Empty 테스트
package dotnet

import "testing"

func TestExtractMapGroups_Empty(t *testing.T) {
	fi := csFileInfo(t, `class C {}`)
	groups := extractMapGroups([]*fileInfo{fi})
	if len(groups) != 0 {
		t.Fatalf("got %v", groups)
	}
}
