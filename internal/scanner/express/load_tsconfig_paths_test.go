//ff:func feature=scan type=test control=sequence topic=express
//ff:what loadTsconfigPaths: paths+baseUrl / baseUrl없음 / 빈targets / 무효json / 파일없음
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

func TestLoadTsconfigPaths_NoBaseUrl(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "tsconfig.json", `{"compilerOptions":{"paths":{"@x/*":["lib/*"]}}}`)
	aliases := loadTsconfigPaths(dir)
	if aliases["@x/"] != "lib/" {
		t.Fatalf("got %q", aliases["@x/"])
	}
}

func TestLoadTsconfigPaths_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "tsconfig.json", `{ not valid`)
	if aliases := loadTsconfigPaths(dir); len(aliases) != 0 {
		t.Fatalf("expected empty for invalid json, got %v", aliases)
	}
}

func TestLoadTsconfigPaths_NoFile(t *testing.T) {
	dir := t.TempDir()
	if aliases := loadTsconfigPaths(dir); len(aliases) != 0 {
		t.Fatalf("expected empty when no tsconfig, got %v", aliases)
	}
}
