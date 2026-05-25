//ff:func feature=sql type=parse control=sequence
//ff:what TestRenderJSON_Empty 테스트
package sqls

import "testing"

func TestRenderJSON_Empty(t *testing.T) {
	result := &SkeletonResult{}
	data, err := RenderJSON(result)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}
