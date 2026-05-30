//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestReadTsconfigJSON_AppFallback 테스트
package express

import "testing"

func TestReadTsconfigJSON_AppFallback(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "tsconfig.app.json", `{"a":1}`)
	if data := readTsconfigJSON(dir); data == nil {
		t.Fatal("expected data from tsconfig.app.json")
	}
}
