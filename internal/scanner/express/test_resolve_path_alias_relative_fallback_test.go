//ff:func feature=scan type=test control=sequence topic=express
//ff:what 상대 경로 import는 path alias 없이도 정상 동작한다
package express

import "testing"

func TestResolvePathAlias_RelativeFallback(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/users.ts", `export const router = express.Router();`)

	got := resolveRelativePath(dir, "./routes/users")
	if got == "" {
		t.Fatal("expected resolved path for relative import, got empty")
	}
}
