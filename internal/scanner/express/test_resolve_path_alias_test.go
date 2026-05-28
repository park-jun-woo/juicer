//ff:func feature=scan type=test control=sequence topic=express
//ff:what tsconfig path alias 해석 테스트: @/api/users/router → src/api/users/router.ts
package express

import "testing"

func TestResolvePathAlias(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/api/users/router.ts", `export const router = express.Router();`)

	aliases := map[string]string{"@/": "src/"}
	got := resolvePathAlias(dir, "@/api/users/router", aliases)
	if got == "" {
		t.Fatal("expected resolved path, got empty")
	}
	want := dir + "/src/api/users/router.ts"
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
