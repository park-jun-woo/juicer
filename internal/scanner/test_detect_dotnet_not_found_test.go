//ff:func feature=scan type=test control=sequence
//ff:what TestDetectDotnet_NotFound -- .csproj 없으면 감지 안 됨
package scanner

import "testing"

func TestDetectDotnet_NotFound(t *testing.T) {
	dir := t.TempDir()
	if detectDotnet(dir) {
		t.Error("expected no dotnet detected")
	}
}
