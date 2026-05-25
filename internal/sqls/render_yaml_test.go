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
