//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestParseFile_Success 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestParseFile_Success(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", `const r = express.Router();`)
	fi, err := parseFile(filepath.Join(dir, "app.ts"))
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if fi == nil || fi.Root == nil {
		t.Fatal("nil fileInfo/root")
	}
}
