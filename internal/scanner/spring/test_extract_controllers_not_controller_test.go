//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractControllers_NotController 테스트
package spring

import "testing"

func TestExtractControllers_NotController(t *testing.T) {
	fi := sFileInfo(t, `public class PlainService {}`)
	if c := extractControllers(fi); c != nil {
		t.Fatalf("expected nil, got %+v", c)
	}
}
