//ff:func feature=scan type=test control=iteration dimension=1
//ff:what newRootCmd 및 하위 명령 생성자가 올바른 Use/하위명령을 구성하는지 직접 테스트
package main

import (
	"strings"
	"testing"
)

func TestNewRootCmd(t *testing.T) {
	root := newRootCmd()
	if !strings.HasPrefix(root.Use, "codist") {
		t.Errorf("root.Use = %q", root.Use)
	}
	want := map[string]bool{"scan": true, "ddl": true, "prisma": true, "sql": true, "version": true}
	got := map[string]bool{}
	for _, c := range root.Commands() {
		got[strings.Fields(c.Use)[0]] = true
	}
	for w := range want {
		if !got[w] {
			t.Errorf("root missing subcommand %q (have %v)", w, got)
		}
	}
}
