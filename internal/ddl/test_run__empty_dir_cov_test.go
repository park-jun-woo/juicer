//ff:func feature=ddl type=command control=sequence
//ff:what TestRun_EmptyDirCov 테스트
package ddl

import "testing"

func TestRun_EmptyDirCov(t *testing.T) {
	dir := t.TempDir()
	out, err := Run(dir)
	if err != nil {
		t.Fatal(err)
	}
	if out != "" {
		t.Fatalf("expected empty, got %q", out)
	}
}
