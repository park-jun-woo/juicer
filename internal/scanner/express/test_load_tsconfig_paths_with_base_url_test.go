//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestLoadTsconfigPaths_WithBaseUrl 테스트
package express

import "testing"

func TestLoadTsconfigPaths_WithBaseUrl(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "tsconfig.json", `{
	  "compilerOptions": {
	    "baseUrl": "src",
	    "paths": { "@app/*": ["app/*"], "@empty/*": [] }
	  }
	}`)
	aliases := loadTsconfigPaths(dir)
	if got, ok := aliases["@app/"]; !ok {
		t.Fatalf("missing @app/ alias: %v", aliases)
	} else if got == "" {
		t.Fatalf("empty replacement")
	}
	if _, ok := aliases["@empty/"]; ok {
		t.Fatalf("empty targets should be skipped: %v", aliases)
	}
}
