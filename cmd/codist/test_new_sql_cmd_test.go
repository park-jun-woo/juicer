//ff:func feature=sql type=test control=iteration dimension=1
//ff:what newSQLCmd Use/플래그/하위명령 등록 직접 테스트
package main

import (
	"strings"
	"testing"
)

func TestNewSQLCmd(t *testing.T) {
	cmd := newSQLCmd()
	if !strings.HasPrefix(cmd.Use, "sql") {
		t.Errorf("Use = %q", cmd.Use)
	}
	for _, name := range []string{"json", "output"} {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("missing --%s flag", name)
		}
	}
	sub := map[string]bool{}
	for _, c := range cmd.Commands() {
		sub[strings.Fields(c.Use)[0]] = true
	}
	for _, w := range []string{"next", "status", "list", "skip", "reset"} {
		if !sub[w] {
			t.Errorf("missing subcommand %q (have %v)", w, sub)
		}
	}
}
