//ff:func feature=sql type=render control=sequence
//ff:what TestRenderYAML_Empty 테스트
package sqls

import "testing"

func TestRenderYAML_Empty(t *testing.T) {
	result := &SkeletonResult{}
	data, err := RenderYAML(result)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}
