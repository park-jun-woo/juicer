//ff:func feature=scan type=test control=sequence topic=express
//ff:what readTsconfigJSON: tsconfig.json / tsconfig.app.json 폴백 / 없음
package express

import "testing"

func TestReadTsconfigJSON_Primary(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "tsconfig.json", `{"a":1}`)
	if data := readTsconfigJSON(dir); data == nil {
		t.Fatal("expected data")
	}
}

func TestReadTsconfigJSON_AppFallback(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "tsconfig.app.json", `{"a":1}`)
	if data := readTsconfigJSON(dir); data == nil {
		t.Fatal("expected data from tsconfig.app.json")
	}
}

func TestReadTsconfigJSON_None(t *testing.T) {
	dir := t.TempDir()
	if data := readTsconfigJSON(dir); data != nil {
		t.Fatalf("expected nil, got %q", data)
	}
}
