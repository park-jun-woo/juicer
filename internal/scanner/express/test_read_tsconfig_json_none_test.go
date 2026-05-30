//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestReadTsconfigJSON_None 테스트
package express

import "testing"

func TestReadTsconfigJSON_None(t *testing.T) {
	dir := t.TempDir()
	if data := readTsconfigJSON(dir); data != nil {
		t.Fatalf("expected nil, got %q", data)
	}
}
