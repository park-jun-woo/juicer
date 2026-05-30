//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestLoadTsconfigPaths_InvalidJSON 테스트
package express

import "testing"

func TestLoadTsconfigPaths_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "tsconfig.json", `{ not valid`)
	if aliases := loadTsconfigPaths(dir); len(aliases) != 0 {
		t.Fatalf("expected empty for invalid json, got %v", aliases)
	}
}
