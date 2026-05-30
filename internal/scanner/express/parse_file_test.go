//ff:func feature=scan type=test control=sequence topic=express
//ff:what parseFile: 정상 파싱 / 읽기 실패
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

func TestParseFile_ReadError(t *testing.T) {
	_, err := parseFile(filepath.Join(t.TempDir(), "missing.ts"))
	if err == nil {
		t.Fatal("expected read error")
	}
}
