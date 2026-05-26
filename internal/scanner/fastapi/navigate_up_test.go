//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what navigateUp 테스트
package fastapi

import "testing"

func TestNavigateUp(t *testing.T) {
	got := navigateUp("/a/b/c", 1)
	if got != "/a/b/c" {
		t.Errorf("dots=1: got %q", got)
	}
	got = navigateUp("/a/b/c", 2)
	if got != "/a/b" {
		t.Errorf("dots=2: got %q", got)
	}
	got = navigateUp("/a/b/c", 3)
	if got != "/a" {
		t.Errorf("dots=3: got %q", got)
	}
}
