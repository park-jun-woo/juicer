//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestMergeMiddleware 테스트
package laravel

import "testing"

func TestMergeMiddleware(t *testing.T) {
	if got := mergeMiddleware([]string{"a"}, nil); len(got) != 1 {
		t.Fatal("empty b")
	}
	got := mergeMiddleware([]string{"a", "b"}, []string{"b", "c"})
	if len(got) != 3 {
		t.Fatalf("dedup failed: %v", got)
	}
}
