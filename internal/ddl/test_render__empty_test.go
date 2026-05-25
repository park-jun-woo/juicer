//ff:func feature=ddl type=render control=sequence
//ff:what TestRender_Empty 테스트
package ddl

import "testing"

func TestRender_Empty(t *testing.T) {
	out := Render(map[string]*Table{})
	if out != "" {
		t.Fatalf("expected empty, got %q", out)
	}
}
