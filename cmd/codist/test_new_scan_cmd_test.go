//ff:func feature=scan type=test control=iteration dimension=1
//ff:what newScanCmd Use 및 플래그 등록 직접 테스트
package main

import (
	"strings"
	"testing"
)

func TestNewScanCmd(t *testing.T) {
	cmd := newScanCmd()
	if !strings.HasPrefix(cmd.Use, "scan") {
		t.Errorf("Use = %q", cmd.Use)
	}
	for _, name := range []string{"json", "openapi", "base", "output", "framework"} {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("missing --%s flag", name)
		}
	}
}
