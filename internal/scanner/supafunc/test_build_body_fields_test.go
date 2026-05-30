//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestBuildBodyFields 테스트
package supafunc

import "testing"

func TestBuildBodyFields(t *testing.T) {
	fields := buildBodyFields([]string{"a", "b"})
	if len(fields) != 2 || fields[0].Name != "a" || fields[0].Type != "unknown" {
		t.Fatalf("got %+v", fields)
	}
}
