//ff:func feature=scan type=test control=sequence
//ff:what TestDetectDotnet_NoCsproj 테스트
package scanner

import "testing"

func TestDetectDotnet_NoCsproj(t *testing.T) {
	if detectDotnet(t.TempDir()) {
		t.Fatal("expected false when no csproj")
	}
}
