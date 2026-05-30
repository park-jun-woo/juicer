//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestReadEnvPrefix 테스트
package nestjs

import "testing"

func TestReadEnvPrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, ".env", "# comment\nAPI_PREFIX=api/v1\nOTHER=x\n")
	if got := readEnvPrefix(dir, ".env"); got != "api/v1" {
		t.Fatalf("got %q", got)
	}
	if got := readEnvPrefix(dir, ".missing"); got != "" {
		t.Fatalf("got %q", got)
	}
}
