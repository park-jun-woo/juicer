//ff:func feature=ddl type=parse control=sequence
//ff:what TestRun_EmptyDir 테스트
package ddl

import (
	"testing"
)

func TestRun_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	got, err := Run(dir)
	if err != nil {
		t.Fatalf("Run() error: %v", err)
	}
	if got != "" {
		t.Errorf("Run() on empty dir should return empty string, got: %q", got)
	}
}
