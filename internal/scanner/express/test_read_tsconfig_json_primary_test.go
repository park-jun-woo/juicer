//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestReadTsconfigJSON_Primary 테스트
package express

import "testing"

func TestReadTsconfigJSON_Primary(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "tsconfig.json", `{"a":1}`)
	if data := readTsconfigJSON(dir); data == nil {
		t.Fatal("expected data")
	}
}
