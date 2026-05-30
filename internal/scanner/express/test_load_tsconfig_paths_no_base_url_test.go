//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestLoadTsconfigPaths_NoBaseUrl 테스트
package express

import "testing"

func TestLoadTsconfigPaths_NoBaseUrl(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "tsconfig.json", `{"compilerOptions":{"paths":{"@x/*":["lib/*"]}}}`)
	aliases := loadTsconfigPaths(dir)
	if aliases["@x/"] != "lib/" {
		t.Fatalf("got %q", aliases["@x/"])
	}
}
